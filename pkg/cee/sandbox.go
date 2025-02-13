package cee

import (
	"github.com/livensmi1e/tiny-ide/pkg/domain"
)

type Sandbox interface {
	Setup(s *domain.Submission)
	Execute(s *domain.Submission) *domain.Metadata
	CleanUp(s *domain.Submission)
	Err() error
}
