package handler

import (
	"github.com/klimenkokayot/avito-go/api_gateway/config"
	"github.com/klimenkokayot/avito-go/api_gateway/internal/domain/ports"
	"github.com/klimenkokayot/avito-go/api_gateway/internal/domain/service"
	"github.com/klimenkokayot/avito-go/libs/logger"
)

type ProxyHandler struct {
	client  ports.AuthService
	service *service.ProxyService
	logger  logger.Logger
	cfg     *config.Config
}

func NewProxyHandler(client ports.AuthService, service *service.ProxyService, logger logger.Logger, cfg *config.Config) (*ProxyHandler, error) {
	return &ProxyHandler{
		client:  client,
		service: service,
		logger:  logger,
		cfg:     cfg,
	}, nil
}
