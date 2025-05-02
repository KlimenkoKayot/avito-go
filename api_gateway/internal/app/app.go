package app

import (
	"context"

	"github.com/klimenkokayot/avito-go/api_gateway/config"
	"github.com/klimenkokayot/avito-go/api_gateway/internal/domain/interfaces"
	"github.com/klimenkokayot/avito-go/api_gateway/internal/domain/service"
	"github.com/klimenkokayot/avito-go/api_gateway/internal/infrastructure/http/client"
	"github.com/klimenkokayot/avito-go/api_gateway/internal/infrastructure/http/server"
	"github.com/klimenkokayot/avito-go/api_gateway/internal/interfaces/http/handler"
	"github.com/klimenkokayot/avito-go/libs/logger"
)

type Application struct {
	server interfaces.ProxyServer
	logger logger.Logger
	cfg    *config.Config
}

// Metrics implements interfaces.Application.
func (a *Application) Metrics() (metrics map[string]interface{}, err error) {
	panic("unimplemented")
}

// Shutdown implements interfaces.Application.
func (a *Application) Shutdown(ctx context.Context) error {
	panic("unimplemented")
}

func (a *Application) Run() error {
	return nil
}

func NewApplication(logger logger.Logger, cfg *config.Config) (interfaces.Application, error) {
	clientLogger := logger.WithLayer("CLIENT")
	client, err := client.NewAuthClient(cfg.AuthPath, clientLogger, cfg)
	if err != nil {
		return nil, err
	}

	serviceLogger := logger.WithLayer("SERVICE")
	service, err := service.NewProxyService(serviceLogger, cfg)
	if err != nil {
		return nil, err
	}

	handlerLogger := logger.WithLayer("HANDLER")
	handler, err := handler.NewProxyHandler(client, service, handlerLogger, cfg)
	if err != nil {
		return nil, err
	}

	serverLogger := logger.WithLayer("SERVER")
	server, err := server.NewProxyServer(handler, serverLogger, cfg)
	if err != nil {
		return nil, err
	}
	return &Application{
		server: server,
		logger: logger,
		cfg:    cfg,
	}, nil
}
