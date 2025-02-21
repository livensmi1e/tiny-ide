package server

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	WORKSPACE       = "sandbox"
	DEFAULT_TIMEOUT = time.Minute
	PYTHON          = "py"
	C               = "c"
	CPP             = "cpp"
	GO              = "go"
)

var languageExtensions = map[int32]string{
	1: PYTHON,
	2: C,
	3: CPP,
	4: GO,
}

type metadata struct {
	languageID int32
	sourceCode string
	filePath   string

	stdout string
	stderr string
}

type engine struct {
	submissionID string
	metadata     metadata
	cmdStrategy  CmdStrategy
	err          error
}

func newEngine(submissionID string, languageID int32, sourceCode string) *engine {
	strategy := getCmdStrategy(languageID)
	if strategy == nil {
		return &engine{err: fmt.Errorf("unsupported language ID: %d", languageID)}
	}
	return &engine{
		submissionID: submissionID,
		metadata: metadata{
			languageID: languageID,
			sourceCode: sourceCode,
		},
		cmdStrategy: strategy,
	}
}

func (e *engine) prepare() {
	if e.err != nil {
		return
	}
	fileExtension, exists := languageExtensions[e.metadata.languageID]
	if !exists {
		e.err = fmt.Errorf("unsupported language ID: %d", e.metadata.languageID)
		return
	}
	e.metadata.filePath, _ = filepath.Abs(filepath.Join(WORKSPACE, fmt.Sprintf("%s.%s", e.submissionID, fileExtension)))
	if err := os.MkdirAll(WORKSPACE, os.ModePerm); err != nil {
		e.err = fmt.Errorf("failed to create workspace: %v", err)
		return
	}
	file, err := os.Create(e.metadata.filePath)
	if err != nil {
		e.err = fmt.Errorf("failed to create file: %v", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(decodeBase64(e.metadata.sourceCode))
	if err != nil {
		e.err = fmt.Errorf("failed to write source code: %v", err)
		return
	}
}

func (e *engine) run() {
	if e.err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), DEFAULT_TIMEOUT)
	defer cancel()
	name, args := e.cmdStrategy.Build(e.metadata.filePath)
	cmd := exec.CommandContext(ctx, name, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		e.err = fmt.Errorf("failed to run source code: %v", err)
		return
	}
	if ctx.Err() == context.DeadlineExceeded {
		e.err = fmt.Errorf("execution timeout")
		return
	}
	e.metadata.stdout = stdout.String()
	e.metadata.stderr = stderr.String()
}

func (e *engine) cleanup() {
	// We don't check err here because we need to clean up file anyway.
	if err := os.RemoveAll(e.metadata.filePath); err != nil {
		e.err = err
	}
	if err := os.RemoveAll(e.cmdStrategy.GetExecutable(e.metadata.filePath)); err != nil {
		e.err = err
	}
}

func (e *engine) error() error {
	return e.err
}

type CmdStrategy interface {
	Build(sourceFile string) (string, []string)
	GetExecutable(sourceFile string) string
}

func getCmdStrategy(l int32) CmdStrategy {
	switch l {
	case 1:
		return &pythonStrategy{}
	case 2:
		return &cStrategy{}
	case 3:
		return &cppStrategy{}
	default:
		return nil
	}
}

type pythonStrategy struct{}

func (p *pythonStrategy) Build(sourceFile string) (string, []string) {
	return "python", []string{sourceFile}
}

func (p *pythonStrategy) GetExecutable(sourceFile string) string {
	return ""
}

type cStrategy struct{}

func (c *cStrategy) Build(sourceFile string) (string, []string) {
	sourceFile = filepath.ToSlash(sourceFile)
	execFile := c.GetExecutable(sourceFile)
	return "sh", []string{"-c", fmt.Sprintf("gcc -o %s %s && %s", execFile, sourceFile, execFile)}
}

func (c *cStrategy) GetExecutable(sourceFile string) string {
	execFile := strings.TrimSuffix(sourceFile, ".c")
	if runtime.GOOS == "windows" {
		execFile += ".exe"
	}
	return execFile
}

type cppStrategy struct{}

func (c *cppStrategy) Build(sourceFile string) (string, []string) {
	sourceFile = filepath.ToSlash(sourceFile)
	execFile := c.GetExecutable(sourceFile)
	return "sh", []string{"-c", fmt.Sprintf("g++ -o %s %s && %s", execFile, sourceFile, execFile)}
}

func (c *cppStrategy) GetExecutable(sourceFile string) string {
	execFile := strings.TrimSuffix(sourceFile, ".cpp")
	if runtime.GOOS == "windows" {
		execFile += ".exe"
	}
	return execFile
}
