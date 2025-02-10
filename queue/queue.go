package queue

type Queue interface {
	Push(submission *Submission) error
	Pop() (*Submission, error)
}
