package service

import (
	"context"
	"net/http"

	"github.com/klimenkokayot/avito-go/api_gateway/config"
	"github.com/klimenkokayot/avito-go/api_gateway/internal/domain/interfaces"
	"github.com/klimenkokayot/avito-go/api_gateway/internal/domain/model"
	"github.com/klimenkokayot/avito-go/libs/logger"
)

type ProxyService struct {
	logger logger.Logger
	cfg    *config.Config
}

// ForwardRequest implements interfaces.ProxyService.
func (p *ProxyService) ForwardRequest(ctx context.Context, r *http.Request) (resp *http.Response, err error) {
	panic("unimplemented")
}

// GetAvailableService implements interfaces.ProxyService.
func (p *ProxyService) GetAvailableService() []string {
	panic("unimplemented")
}

// HandleProxyError implements interfaces.ProxyService.
func (p *ProxyService) HandleProxyError(w http.ResponseWriter, err error) error {
	panic("unimplemented")
}

// Metrics implements interfaces.ProxyService.
func (p *ProxyService) Metrics() (metrics map[string]interface{}, err error) {
	panic("unimplemented")
}

// RegisterService implements interfaces.ProxyService.
func (p *ProxyService) RegisterService(name string, config *model.ServiceConfig) error {
	panic("unimplemented")
}

func NewProxyService(logger logger.Logger, cfg *config.Config) (interfaces.ProxyService, error) {
	return &ProxyService{
		logger: logger,
		cfg:    cfg,
	}, nil
}
