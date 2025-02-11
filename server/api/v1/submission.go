package v1

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/livensmi1e/tiny-ide/pkg/domain"
	"github.com/livensmi1e/tiny-ide/pkg/wrapper"
	"github.com/livensmi1e/tiny-ide/server/api/v1/dto"
	"github.com/livensmi1e/tiny-ide/store"
)

func (w *WebAPIV1) HandleSubmission(c echo.Context) *wrapper.Response {
	sub := new(dto.Submission)
	if err := c.Bind(&sub); err != nil {
		return &wrapper.Response{
			Code:  http.StatusBadRequest,
			Error: fmt.Errorf("invalid request body: %w", err).Error(),
		}
	}
	if err := w.infra.Validator().Validate(sub); err != nil {
		return &wrapper.Response{
			Code:  http.StatusBadRequest,
			Error: fmt.Errorf("failed to parse request: %w", err).Error(),
		}
	}
	subDB, err := w.infra.Store().CreateSubmission(&store.Submission{
		ID:         randomString(6),
		LanguageID: sub.LanguageID,
	})
	if err != nil {
		return &wrapper.Response{
			Code:  http.StatusInternalServerError,
			Error: fmt.Errorf("failed to create db entry: %w", err).Error(),
		}
	}
	w.infra.Queue().Push(&domain.Submission{
		ID:         subDB.ID,
		LanguageID: sub.LanguageID,
		SourceCode: sub.SourceCode,
	})
	return &wrapper.Response{
		Code: http.StatusCreated,
		Data: subDB.ID,
	}
}

func (w *WebAPIV1) GetSubmission(c echo.Context) *wrapper.Response {
	token := c.Param("token")
	subDB, err := w.infra.Store().GetSubmission(&store.FindSubmission{
		ID:    &token,
		Limit: nil,
	})
	if err != nil {
		return &wrapper.Response{
			Code:  http.StatusInternalServerError,
			Error: fmt.Errorf("failed to get db entry: %w", err).Error(),
		}
	}
	return &wrapper.Response{
		Code: http.StatusOK,
		Data: convertFromStore(subDB),
	}
}

func randomString(length int) string {
	if length == 0 {
		length = 6
	}
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func convertFromStore(sub *store.Submission) *dto.SubmissionResponse {
	return &dto.SubmissionResponse{
		ID:         sub.ID,
		LanguageID: sub.LanguageID,
		Status:     sub.Status,
		Stdout:     sub.Stdout,
		Stderr:     sub.Stderr,
		Time:       sub.Time,
		Memory:     sub.Memory,
	}
}
