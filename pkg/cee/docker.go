package cee

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"time"

	"github.com/livensmi1e/tiny-ide/pkg/constant"
	"github.com/livensmi1e/tiny-ide/pkg/domain"
)

type dockerContainer struct {
	ImageName string
	Timeout   time.Duration
}

func NewDockerContainer(i string, t time.Duration) *dockerContainer {
	return &dockerContainer{
		ImageName: i,
		Timeout:   t,
	}
}

func (d *dockerContainer) BuildCommand(lang string) string {
	switch lang {
	case constant.PYTHON:
		return "/usr/bin/time -v python3 -c"
	case constant.C:
		return "/usr/bin/time -v sh -c 'gcc -x c -o /tmp/result.out - && /tmp/result.out'"
	case constant.CPP:
		return "/usr/bin/time -v sh -c 'g++ -x c++ -o /tmp/result.out - && /tmp/result.out'"
	default:
		return ""
	}
}

func (d *dockerContainer) Run(s *domain.Submission) (*domain.Metadata, error) {
	execCmd := d.BuildCommand(s.MapLang())
	if execCmd == "" {
		return &domain.Metadata{Stdout: "", Stderr: "", Time: constant.DefaultTime, Memory: constant.DefaultMemory}, fmt.Errorf("unsupported language: %s", s.MapLang())
	}
	execCmd += " " + "'" + s.SourceCode + "'"

	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout)
	defer cancel()

	runCmd := exec.CommandContext(ctx, "docker", "run", "--rm", "-i", d.ImageName, "sh", "-c", execCmd)
	var stdout, stderr bytes.Buffer
	runCmd.Stdout = &stdout
	runCmd.Stderr = &stderr

	if err := runCmd.Run(); err != nil {
		return &domain.Metadata{Stdout: "", Stderr: removeStderrStats(stderr.String()), Time: constant.DefaultTime, Memory: constant.DefaultMemory}, fmt.Errorf("run container failed %w", err)
	}

	if ctx.Err() == context.DeadlineExceeded {
		return &domain.Metadata{Stdout: "", Stderr: "execution timeout", Time: constant.DefaultTime, Memory: constant.DefaultMemory}, fmt.Errorf("execution timeout")
	}

	time, memory := extractTimeAndMemory(stderr.String())

	return &domain.Metadata{Stdout: stdout.String(), Stderr: removeStderrStats(stderr.String()), Time: convertTime(time), Memory: memory}, nil
}

func (d *dockerContainer) Clean() error {
	return nil
}

func extractTimeAndMemory(stderr string) (string, string) {
	reTime := regexp.MustCompile(`Elapsed \(wall clock\) time \(h:mm:ss or m:ss\): ([^\n]+)`)
	reMemory := regexp.MustCompile(`Maximum resident set size \(kbytes\): (\d+)`)
	timeMatch := reTime.FindStringSubmatch(stderr)
	memoryMatch := reMemory.FindStringSubmatch(stderr)
	timeResult := constant.DefaultTime
	memoryResult := constant.DefaultMemory
	if len(timeMatch) > 1 {
		timeResult = timeMatch[1]
	}
	if len(memoryMatch) > 1 {
		memoryResult = memoryMatch[1]
	}
	return timeResult, memoryResult
}

func removeStderrStats(stderr string) string {
	re := regexp.MustCompile(`(?s)(.*?)\n?\tCommand being timed:`)
	match := re.FindStringSubmatch(stderr)
	if len(match) > 1 {
		return match[1]
	}
	return stderr
}

func convertTime(time string) string {
	re := regexp.MustCompile(`(\d+):(\d{2}\.\d{2})`)
	match := re.FindStringSubmatch(time)
	if len(match) != 3 {
		return constant.DefaultTime
	}
	min, _ := strconv.Atoi(match[1])
	sec, _ := strconv.ParseFloat(match[2], 64)
	totalSec := float64(min*60) + sec
	return fmt.Sprintf("%.4f", totalSec)
}
