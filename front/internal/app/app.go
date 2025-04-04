package app

import (
	"github.com/klimenkokayot/avito-go/front/config"
	server "github.com/klimenkokayot/avito-go/front/internal/infrastructure/http"
	"github.com/klimenkokayot/avito-go/front/internal/interfaces/http/handlers"
	"github.com/klimenkokayot/avito-go/libs/logger"
)

type Application struct {
	server *server.Server
	logger logger.Logger
	cfg    *config.Config
}

func NewApplication(cfg *config.Config, logger logger.Logger) (*Application, error) {
	handlersLogger := logger.WithLayer("HANDLERS")
	handlers, err := handlers.NewViewHandler(cfg, handlersLogger)
	if err != nil {
		return nil, err
	}

	serverLogger := logger.WithLayer("SERVER")
	server, err := server.NewServer(handlers, cfg, serverLogger)
	if err != nil {
		return nil, err
	}
	return &Application{
		server: server,
		logger: logger,
		cfg:    cfg,
	}, nil
}

func (a *Application) Run() error {
	return a.server.Run()
}
