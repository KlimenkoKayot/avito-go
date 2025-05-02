package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"path"

	"github.com/klimenkokayot/avito-go/api_gateway/config"
	"github.com/klimenkokayot/avito-go/api_gateway/internal/domain/model"
	"github.com/klimenkokayot/avito-go/api_gateway/internal/domain/ports"
	"github.com/klimenkokayot/avito-go/libs/logger"
)

type AuthClient struct {
	authBaseURL string
	client      *http.Client
	logger      logger.Logger
	cfg         *config.Config
}

// Metrics implements ports.AuthService.
func (a *AuthClient) Metrics() (metrics map[string]interface{}, err error) {
	panic("unimplemented")
}

// RefreshTokens implements ports.AuthService.
func (a *AuthClient) RefreshTokens(ctx context.Context, refreshToken string) (tokenPair *model.TokenPair, err error) {
	panic("unimplemented")
}

// Проверяет пару токенов
func (a *AuthClient) Authenticate(ctx context.Context, tokenPair *model.TokenPair) (userID string, err error) {
	a.logger.Info("Проверка токенов.")
	data, err := json.Marshal(tokenPair)
	if err != nil {
		a.logger.Warn("Не удалось сформировать тело запроса.")
		return "", err
	}
	buffer := new(bytes.Buffer)
	buffer.WriteString(string(data))
	req, err := http.NewRequest(http.MethodPost, path.Join(a.authBaseURL, "auth", "validate"), buffer)
	if err != nil {
		a.logger.Warn("Не удалось сформировать http.Request.")
		return "", err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		a.logger.Warn("Не удалось отравить запрос.")
		return "", err
	}
	defer resp.Body.Close()

	// Проверка на верификацию
	if resp.StatusCode == http.StatusOK {
		a.logger.OK("Токен валиден.")
		return "todo_user_id", nil
	}
	a.logger.OK("Токен не валиден.")
	return "", nil
}

func NewAuthClient(authBaseURL string, logger logger.Logger, cfg *config.Config) (ports.AuthService, error) {
	logger.Info("Создание AuthClient`а.")
	logger.OK("Успешное создание AuthClient`a.")
	return &AuthClient{
		authBaseURL: authBaseURL,
		client:      &http.Client{},
		logger:      logger,
		cfg:         cfg,
	}, nil
}
