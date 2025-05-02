package server

import (
	"context"
	"net/http"

	"github.com/klimenkokayot/avito-go/api_gateway/config"
	"github.com/klimenkokayot/avito-go/api_gateway/internal/domain/interfaces"
	"github.com/klimenkokayot/avito-go/libs/logger"
	"github.com/klimenkokayot/avito-go/libs/router"
)

type ProxyServer struct {
	router  router.Router
	handler interfaces.ProxyHandler
	logger  logger.Logger
	cfg     *config.Config
}

// Metrics implements interfaces.ProxyServer.
func (p *ProxyServer) Metrics() (metrics map[string]interface{}, err error) {
	panic("unimplemented")
}

// RateLimitMiddleware implements interfaces.ProxyServer.
func (p *ProxyServer) RateLimitMiddleware(next http.Handler) http.Handler {
	panic("unimplemented")
}

// RegisterMiddleware implements interfaces.ProxyServer.
func (p *ProxyServer) RegisterMiddleware(middleware func(http.Handler) http.Handler) error {
	panic("unimplemented")
}

// RegisterRoute implements interfaces.ProxyServer.
func (p *ProxyServer) RegisterRoute(method string, path string, route http.HandlerFunc) error {
	panic("unimplemented")
}

// Run implements interfaces.ProxyServer.
func (p *ProxyServer) Run() error {
	panic("unimplemented")
}

// Shutdown implements interfaces.ProxyServer.
func (p *ProxyServer) Shutdown(ctx context.Context) error {
	panic("unimplemented")
}

func NewProxyServer(handler interfaces.ProxyHandler, logger logger.Logger, cfg *config.Config) (interfaces.ProxyServer, error) {
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
