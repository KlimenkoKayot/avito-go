package server

import (
	"github.com/klimenkokayot/avito-go/api_gateway/config"
	"github.com/klimenkokayot/avito-go/api_gateway/internal/interfaces/http/handler"
	"github.com/klimenkokayot/avito-go/libs/logger"
	"github.com/klimenkokayot/avito-go/libs/router"
)

type ProxyServer struct {
	router  router.Router
	handler *handler.ProxyHandler
	logger  logger.Logger
	cfg     *config.Config
}

func NewProxyServer(handler *handler.ProxyHandler, logger logger.Logger, cfg *config.Config) (*ProxyServer, error) {
	router, err := router.NewAdapter(&router.Config{
		Name: cfg.Router,
	})
	if err != nil {
		return nil, err
	}
	return &ProxyServer{
		router:  router,
		handler: handler,
		logger:  logger,
		cfg:     cfg,
	}, nil
}
