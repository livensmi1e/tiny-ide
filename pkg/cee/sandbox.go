package cee

import (
	"github.com/livensmi1e/tiny-ide/pkg/domain"
)

type Sandbox interface {
	Run(s *domain.Submission) (*domain.Metadata, error)
	BuildCommand(language string, sourceCode string) string
	Clean() error
}
