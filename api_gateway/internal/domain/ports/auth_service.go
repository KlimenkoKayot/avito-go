package ports

import (
	"context"

	"github.com/klimenkokayot/avito-go/api_gateway/internal/domain/model"
)

type AuthService interface {
	// Аутентификация
	Authenticate(ctx context.Context, tokenPair *model.TokenPair) (userID string, err error)

	// Обновление токенов
	RefreshTokens(ctx context.Context, refreshToken string) (tokenPair *model.TokenPair, err error)
}
