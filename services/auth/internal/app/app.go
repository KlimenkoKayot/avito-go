package app

import (
	"github.com/klimenkokayot/avito-go/libs/logger"
	"github.com/klimenkokayot/avito-go/services/auth/config"
	"github.com/klimenkokayot/avito-go/services/auth/internal/domain"
	repo "github.com/klimenkokayot/avito-go/services/auth/internal/infrastructure/http/repository"
)

type Application struct {
	config *config.Config
	logger *logger.Logger
	server *domain.Server
}

func NewApplication(cfg *config.Config, logger logger.Logger) (domain.Application, error) {
	repo, err := repo.NewUserRepository(cfg, logger)
	if err != nil {
		return nil, err
	}
	service, err := service.NewService(cfg, logger, repo)
	return &Application{}, nil
}

func (a *Application) Run() error {

	return nil
}
