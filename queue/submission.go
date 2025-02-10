package queue

import "encoding/json"

type Submission struct {
	ID         string `json:"id"`
	LanguageID string `json:"language_id"`
	SourceCode string `json:"source_code"`
}

func (s *Submission) Serialize() (string, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func Deserialize(data string) (*Submission, error) {
	var s Submission
	err := json.Unmarshal([]byte(data), &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
