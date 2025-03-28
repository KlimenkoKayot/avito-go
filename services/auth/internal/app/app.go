package app

import (
	"github.com/klimenkokayot/avito-go/services/auth/config"
	"github.com/klimenkokayot/avito-go/services/auth/internal/domain"
)

type Application struct {
	config *config.Config
	server *domain.Server
}

func NewApplication(cfg *config.Config, logger domain.Logger) (*Application, error) {

}
