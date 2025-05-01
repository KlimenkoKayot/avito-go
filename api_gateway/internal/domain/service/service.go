package service

import (
	"github.com/klimenkokayot/avito-go/api_gateway/config"
	"github.com/klimenkokayot/avito-go/libs/logger"
)

type ProxyService struct {
	logger logger.Logger
	cfg    *config.Config
}

func NewProxyService(logger logger.Logger, cfg *config.Config) (*ProxyService, error) {
	return &ProxyService{
		logger: logger,
		cfg:    cfg,
	}, nil
}
