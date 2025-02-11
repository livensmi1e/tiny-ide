package domain

import (
	"encoding/json"

	"github.com/livensmi1e/tiny-ide/pkg/constant"
)

type Submission struct {
	ID         string `json:"id"`
	LanguageID int32  `json:"language_id"`
	SourceCode string `json:"source_code"`
}

func (s *Submission) Serialize() (string, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (s *Submission) MapLang() string {
	langMap := map[int32]string{
		1: constant.PYTHON,
		2: constant.C,
		3: constant.CPP,
	}
	return langMap[s.LanguageID]
}

func Deserialize(data string) (*Submission, error) {
	var s Submission
	err := json.Unmarshal([]byte(data), &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
