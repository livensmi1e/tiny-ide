package queue

import "github.com/livensmi1e/tiny-ide/pkg/domain"

type Queue interface {
	Push(submission *domain.Submission) error
	Pop() (*domain.Submission, error)
}
