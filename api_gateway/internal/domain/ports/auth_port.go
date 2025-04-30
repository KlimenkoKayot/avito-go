package ports

import (
	"context"

	model "github.com/klimenkokayot/avito-go/api_gateway/internal/domain/model/auth"
)

// Проверяет токен, в случае проблем - redirect в микросервис авторизации.
type AuthService interface {
	VerifyTokenPair(ctx context.Context, tokenPair *model.TokenPair) (userID string, err error)
}
