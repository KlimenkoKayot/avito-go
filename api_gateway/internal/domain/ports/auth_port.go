package ports

import (
	"context"
	"net/http"

	model "github.com/klimenkokayot/avito-go/api_gateway/internal/domain/model/auth"
)

type AuthService interface {
	ForwardAuthRequest(ctx context.Context, r *http.Request) (*http.Response, error)
	VerifyTokenPair(ctx context.Context, tokenPair *model.TokenPair) (bool, error)
	RefreshTokenPair(ctx context.Context, refreshToken string) (*model.TokenPair, error)
	Logout(ctx context.Context, accessToken string) error // aka инвалидация токенов
}
