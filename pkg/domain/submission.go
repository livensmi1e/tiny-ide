package domain

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/livensmi1e/tiny-ide/pkg/constant"
)

type Submission struct {
	ID         string `json:"id"`
	LanguageID int32  `json:"language_id"`
	SourceCode string `json:"source_code"`
	FilePath   string `json:"file_path"`
	FileName   string `json:"file_name"`
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

func (s *Submission) DecodeSource() string {
	decodeByte, _ := base64.StdEncoding.DecodeString(s.SourceCode)
	return string(decodeByte)
}

func (s *Submission) SaveSourceToFile(dirName string) error {
	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create workspace: %v", err)
	}
	sourceCode := s.DecodeSource()
	fileName := fmt.Sprintf("%s.%s", s.ID, s.MapLang())
	filePath := filepath.Join(dirName, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()
	_, err = file.WriteString(sourceCode)
	if err != nil {
		return fmt.Errorf("failed to write source code: %v", err)
	}
	s.FilePath = filePath
	s.FileName = fileName
	return nil
}

func Deserialize(data string) (*Submission, error) {
	var s Submission
	err := json.Unmarshal([]byte(data), &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
