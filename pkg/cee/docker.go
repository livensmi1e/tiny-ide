package cee

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"time"

	"github.com/livensmi1e/tiny-ide/pkg/domain"
)

type dockerContainer struct {
	ImageName string
	Timeout   time.Duration
	err       error
}

func NewDockerContainer(i string, t time.Duration) *dockerContainer {
	return &dockerContainer{
		ImageName: i,
		Timeout:   t,
		err:       nil,
	}
}

func (d *dockerContainer) Setup(s *domain.Submission) {
	d.err = s.SaveSourceToFile("workspace")
}

func (d *dockerContainer) Execute(s *domain.Submission) *domain.Metadata {
	if d.err != nil {
		return &domain.DefaultMetadata
	}

	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout)
	defer cancel()

	runCmd := exec.CommandContext(ctx, "docker", "run",
		"--rm",
		"-v", fmt.Sprintf("%s:/sandbox/code/%s", s.FilePath, s.FileName),
		d.ImageName,
		fmt.Sprintf("/sandbox/code/%s", s.FileName),
	)
	var stdout bytes.Buffer
	runCmd.Stdout = &stdout

	if err := runCmd.Run(); err != nil {
		str := runCmd.String()
		fmt.Println(str)
		d.err = fmt.Errorf("run container failed %w", err)
		return &domain.DefaultMetadata
	}
	if ctx.Err() == context.DeadlineExceeded {
		d.err = fmt.Errorf("execution timeout")
		return &domain.DefaultMetadata
	}

	return d.parseMetadata(stdout.String())
}

func (d *dockerContainer) CleanUp(s *domain.Submission) {
	d.err = s.DeleteFile("workspace")
}

func (d *dockerContainer) Err() error {
	return d.err
}

func (d *dockerContainer) parseMetadata(metadata string) *domain.Metadata {
	parsedMetadata := &domain.DefaultMetadata
	stdoutRegex := regexp.MustCompile(`(?s)stdout:\s*(.*?)\nstderr`)
	stderrRegex := regexp.MustCompile(`(?s)stderr:\s*(.*?)\ntime`)
	timeRegex := regexp.MustCompile(`time:\s*(\d+)\s*ms`)
	memoryRegex := regexp.MustCompile(`memory:\s*(\d+)\s*kb`)
	if match := stdoutRegex.FindStringSubmatch(metadata); match != nil {
		parsedMetadata.Stdout = match[1]
	}
	if match := stderrRegex.FindStringSubmatch(metadata); match != nil {
		parsedMetadata.Stderr = match[1]
	}
	if match := timeRegex.FindStringSubmatch(metadata); match != nil {
		parsedMetadata.Time = match[1]
	}
	if match := memoryRegex.FindStringSubmatch(metadata); match != nil {
		parsedMetadata.Memory = match[1]
	}
	return parsedMetadata
}
