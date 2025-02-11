package store

import "database/sql"

type Driver interface {
	Close() error
	GetDB() *sql.DB
	Migrate() error

	CreateSubmission(create *Submission) (*Submission, error)
	UpdateSubmission(update *UpdateSubmission) (*Submission, error)
	ListSubmissions(find *FindSubmission) ([]*Submission, error)
}
