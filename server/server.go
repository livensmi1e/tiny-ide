package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/livensmi1e/tiny-ide/infra"
	v1 "github.com/livensmi1e/tiny-ide/server/api/v1"
)

type Server struct {
	api    *v1.WebAPIV1
	router *echo.Echo
	infra  infra.Infrastructure
}

func NewServer(infra infra.Infrastructure) *Server {
	server := &Server{
		infra:  infra,
		router: echo.New(),
		api:    v1.New(infra),
	}
	server.setupMiddlewares()
	server.registerRoutes()
	return server
}

func (s *Server) Start() error {
	return s.router.Start(s.infra.Config().Addr + ":" + s.infra.Config().Port)
}

func (s *Server) setupMiddlewares() {
	s.router.Use(middleware.CORS())

	logCfg := s.infra.Config().GetEchoLogConfig()
	s.router.Use(middleware.LoggerWithConfig(logCfg))
}

func (s *Server) registerRoutes() {
	s.api.RegisterHandlers(s.router)
}
