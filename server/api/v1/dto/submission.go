package dto

type Submission struct {
	SourceCode string `json:"source_code"`
	LanguageID int32  `json:"language_id"`
}
