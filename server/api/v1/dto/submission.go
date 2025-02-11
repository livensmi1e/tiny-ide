package dto

type Submission struct {
	SourceCode string `json:"source_code" validate:"required"`
	LanguageID int32  `json:"language_id" validate:"required"`
}

type SubmissionResponse struct {
	ID         string `json:"id"`
	LanguageID int32  `json:"language_id"`
	Status     string `json:"status"`
	Stdout     string `json:"stdout"`
	Stderr     string `json:"stderr"`
	Time       string `json:"time"`
	Memory     string `json:"memory"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
