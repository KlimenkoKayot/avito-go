package app

import (
	"github.com/klimenkokayot/avito-go/libs/logger"
	"github.com/klimenkokayot/avito-go/services/auth/config"
	"github.com/klimenkokayot/avito-go/services/auth/internal/domain"
)

type Application struct {
	config *config.Config
	logger *logger.Logger
	server *domain.Server
}

func NewApplication(cfg *config.Config, logger logger.Logger) (domain.Application, error) {
	return &Application{}, nil
}

func (a *Application) Run() error {
	return nil
}
