package ports

import (
	"context"
	"net/http"

	model "github.com/klimenkokayot/avito-go/api_gateway/internal/domain/model/auth"
)

// Проверяет токен, в случае проблем - проксирует в микросервис авторизации.
type AuthService interface {
	VerifyTokenPair(ctx context.Context, tokenPair *model.TokenPair) (userID string, err error)
	ProxyAuthRequest(ctx context.Context, r *http.Request) (*http.Response, error)
}
