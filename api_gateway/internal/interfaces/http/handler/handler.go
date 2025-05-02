package handler

import (
	"net/http"

	"github.com/klimenkokayot/avito-go/api_gateway/config"
	"github.com/klimenkokayot/avito-go/api_gateway/internal/domain/interfaces"
	"github.com/klimenkokayot/avito-go/api_gateway/internal/domain/ports"
	"github.com/klimenkokayot/avito-go/libs/logger"
)

type ProxyHandler struct {
	client  ports.AuthService
	service interfaces.ProxyService
	logger  logger.Logger
	cfg     *config.Config
}

// AuthMiddleware implements interfaces.ProxyHandler.
func (p *ProxyHandler) AuthMiddleware(next http.Handler) http.Handler {
	panic("unimplemented")
}

// LoggerMiddleware implements interfaces.ProxyHandler.
func (p *ProxyHandler) LoggerMiddleware(next http.Handler) http.Handler {
	panic("unimplemented")
}

// Metrics implements interfaces.ProxyHandler.
func (p *ProxyHandler) Metrics() (metrics map[string]interface{}, err error) {
	panic("unimplemented")
}

// ProxyAuthService implements interfaces.ProxyHandler.
func (p *ProxyHandler) ProxyAuthService(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// ServeHTTP implements interfaces.ProxyHandler.
func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func NewProxyHandler(client ports.AuthService, service interfaces.ProxyService, logger logger.Logger, cfg *config.Config) (interfaces.ProxyHandler, error) {
	return &ProxyHandler{
		client:  client,
		service: service,
		logger:  logger,
		cfg:     cfg,
	}, nil
}
