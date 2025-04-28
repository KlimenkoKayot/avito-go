package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"path"

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
	data, err := json.Marshal(tokenPair)
	if err != nil {
		return "", err
	}
	buffer := new(bytes.Buffer)
	buffer.WriteString(string(data))
	req, err := http.NewRequest(http.MethodGet, path.Join(a.authBaseURL, "auth", "validate"), buffer)
	if err != nil {
		return "", err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Проверка на верификацию
	if resp.StatusCode == http.StatusOK {
		return "todo_user_id", nil
	}
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
