package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/livensmi1e/tiny-ide/infra"
	"github.com/livensmi1e/tiny-ide/pkg/wrapper"
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
	v1.GET("/demo", wrapper.Wrap(w.Demo))

	submission := v1.Group("/submissions")
	{
		submission.POST("", wrapper.Wrap(w.HandleSubmission))
		submission.GET("/:token", wrapper.Wrap(w.GetSubmission))
	}
}

func (w *WebAPIV1) Demo(c echo.Context) *wrapper.Response {
	return &wrapper.Response{
		Code: http.StatusOK,
		Data: "Hello World",
	}
}
