package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/livensmi1e/tiny-ide/infra"
)

type WebAPIV1 struct {
	infra infra.Infrastructure
}

func New(infra infra.Infrastructure) *WebAPIV1 {
	return &WebAPIV1{
		infra: infra,
	}
}

func (w *WebAPIV1) RegisterHandlers(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	submission := v1.Group("/submissions")
	{
		submission.POST("", nil)
		submission.GET("/:token", nil)
	}
}
