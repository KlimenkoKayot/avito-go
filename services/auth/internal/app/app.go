package app

import (
	"github.com/klimenkokayot/avito-go/libs/logger"
	"github.com/klimenkokayot/avito-go/services/auth/config"
	"github.com/klimenkokayot/avito-go/services/auth/internal/domain"
	"github.com/klimenkokayot/avito-go/services/auth/internal/domain/service"
	server "github.com/klimenkokayot/avito-go/services/auth/internal/infrastructure/http"
	repo "github.com/klimenkokayot/avito-go/services/auth/internal/infrastructure/http/repository"
	"github.com/klimenkokayot/avito-go/services/auth/internal/interfaces/http/handlers"
)

type Application struct {
	server domain.Server
	logger logger.Logger
	config *config.Config
}

func NewApplication(cfg *config.Config, logger logger.Logger) (domain.Application, error) {
	repo, err := repo.NewUserRepository(cfg, logger)
	if err != nil {
		return nil, err
	}

	service, err := service.NewAuthService(repo, cfg, logger)
	if err != nil {
		return nil, err
	}

	handler, err := handlers.NewAuthHandler(service, cfg, logger)
	if err != nil {
		return nil, err
	}

	server, err := server.NewAuthServer(handler, cfg, logger)
	if err != nil {
		return nil, err
	}

	return &Application{
		server,
		logger,
		cfg,
	}, nil
}

func (a *Application) Run() error {
	return a.server.Run()
}
