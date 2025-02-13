package cee

import (
	"github.com/livensmi1e/tiny-ide/pkg/domain"
)

// TODO: Remove build command, this job belongs to script .sh inside container
// TODO: Add Prepare and CleanUp method: Save and remove neccessary file
type Sandbox interface {
	Setup(s *domain.Submission)
	Execute(s *domain.Submission) *domain.Metadata
	CleanUp(s *domain.Submission)
	Err() error
}
