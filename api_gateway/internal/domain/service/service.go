package service

import (
	"github.com/klimenkokayot/avito-go/api_gateway/config"
	"github.com/klimenkokayot/avito-go/libs/logger"
)

type ProxyService struct {
	logger logger.Logger
	cfg    *config.Config
}

func (ps *ProxyService) EndpointToURL(endpoint string) (string, error) {
	ps.logger.Warn(endpoint)
	return endpoint, nil
}

func NewProxyService(logger logger.Logger, cfg *config.Config) (*ProxyService, error) {
	return &ProxyService{
		logger: logger,
		cfg:    cfg,
	}, nil
}
