package client

import (
	"context"
	"net/http"

	"github.com/klimenkokayot/avito-go/api_gateway/config"
	model "github.com/klimenkokayot/avito-go/api_gateway/internal/domain/model/auth"
	"github.com/klimenkokayot/avito-go/api_gateway/internal/domain/ports"
	"github.com/klimenkokayot/avito-go/libs/logger"
)

type AuthClient struct {
	client      *http.Client
	authBaseURL string
	logger      *logger.Logger
	cfg         *config.Config
}

// Проксирует запросы в микросервис авторизации
func (a *AuthClient) ProxyAuthRequest(ctx context.Context, r *http.Request) (*http.Response, error) {
	return nil, nil
}

// Проверяет пару токенов
func (a *AuthClient) VerifyTokenPair(ctx context.Context, tokenPair *model.TokenPair) (userID string, err error) {
	return "", nil
}

func NewAuthClient(authBaseURL string, logger *logger.Logger, cfg *config.Config) (ports.AuthService, error) {
	return &AuthClient{
		client:      &http.Client{},
		authBaseURL: authBaseURL,
		logger:      logger,
		cfg:         cfg,
	}, nil
}
