package server

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/klimenkokayot/avito-go/front/config"
	"github.com/klimenkokayot/avito-go/front/internal/infrastructure/http/middleware"
	"github.com/klimenkokayot/avito-go/front/internal/interfaces/http/handlers"
	"github.com/klimenkokayot/avito-go/libs/logger"
	"github.com/klimenkokayot/avito-go/libs/router"
)

type Server struct {
	handlers *handlers.ViewHandler
	router   router.Router
	logger   logger.Logger
	cfg      *config.Config
}

func NewServer(handlers *handlers.ViewHandler, cfg *config.Config, logger logger.Logger) (*Server, error) {
	router, err := router.NewAdapter(&router.Config{
		Name: cfg.Router,
	})
	if err != nil {
		return nil, err
	}

	server := &Server{
		handlers: handlers,
		router:   router,
		logger:   logger,
		cfg:      cfg,
	}

	err = server.setupStatic()
	if err != nil {
		return nil, err
	}

	err = server.setupRoutes()
	if err != nil {
		return nil, err
	}

	err = server.setupMiddlewares()
	if err != nil {
		return nil, err
	}

	return server, nil
}

// TODO FIX
func (s *Server) setupStatic() error {
	staticDir, err := filepath.Abs(filepath.Join("web", "static"))
	if err != nil {
		return err
	}

	// Gorilla Mux требует явного указания /* для подпутей
	fs := http.FileServer(http.Dir(staticDir))
	s.router.Handle("/static/{rest:.*}", http.StripPrefix("/static/", fs))
	return nil
}

func (s *Server) setupRoutes() error {
	s.router.GET("/login", s.handlers.LoginPage)
	s.router.GET("/register", s.handlers.RegisterPage)
	return nil
}

func (s *Server) setupMiddlewares() error {
	s.router.Use(middleware.LoggerMiddleware(s.logger.WithLayer("LOG_MWARE")))
	s.router.Use(middleware.TimeoutMiddleware(s.cfg.ReadTimeoutSeconds, s.cfg.WriteTimeoutSeconds))
	return nil
}

func (s *Server) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.ServerPort), s.router)
}
