package store

type (
	Submission struct {
		ID         string
		LanguageID int32
		Status     string
		Stdout     string
		Stderr     string
		Time       string
		Memory     string
	}
	UpdateSubmission struct {
		ID         string
		LanguageID int32
		Status     *string
		Stdout     *string
		Stderr     *string
		Time       *string
		Memory     *string
	}
	FindSubmission struct {
		ID    *string
		Limit *int
	}
)

func (s *Store) CreateSubmission(create *Submission) (*Submission, error) {
	return s.driver.CreateSubmission(create)
}

func (s *Store) UpdateSubmission(update *UpdateSubmission) (*Submission, error) {
	return s.driver.UpdateSubmission(update)
}

func (s *Store) ListSubmissions(find *FindSubmission) ([]*Submission, error) {
	return s.driver.ListSubmissions(find)
}

func (s *Store) GetSubmission(find *FindSubmission) (*Submission, error) {
	list, err := s.driver.ListSubmissions(find)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}
	return list[0], nil
}
