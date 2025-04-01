package server

import (
	"github.com/klimenkokayot/avito-go/front/config"
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

	return &Server{
		handlers: handlers,
		router:   router,
		logger:   logger,
		cfg:      cfg,
	}, nil
}
