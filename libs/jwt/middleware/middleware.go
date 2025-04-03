package jwt

import (
	"net/http"

	"github.com/klimenkokayot/avito-go/libs/jwt"
)

type TokenMiddleware struct {
	tokenManager *jwt.TokenManager
}

func NewTokenMiddleware(tokenManager *jwt.TokenManager) (*TokenMiddleware, error) {
	return &TokenMiddleware{
		tokenManager: tokenManager,
	}, nil
}

func (a *TokenMiddleware) refreshTokenPair(r *http.Request) error {
	refreshTokenCookie, err := r.Cookie("refresh_token")
	refreshTokenString := refreshTokenCookie.String()

	claims, err := a.tokenManager.ParseWithClaims(refreshTokenString)
	if err != nil {
		return err
	}

	login := (*claims)["lgn"].(string)
	ip := r.RemoteAddr

	newAccessTokenString, err := a.tokenManager.NewAccessToken(login, ip)
	if err != nil {
		return err
	}
	newRefreshTokenString, err := a.tokenManager.NewRefreshToken(login)
	if err != nil {
		return err
	}

	r.AddCookie(&http.Cookie{
		Name:  "access_token",
		Value: newAccessTokenString,
	})
	r.AddCookie(&http.Cookie{
		Name:  "refresh_token",
		Value: newRefreshTokenString,
	})
	return nil
}

func (a *TokenMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessTokenCookie, err := r.Cookie("access_token")
		if err != nil {
			http.Error(w, "access_token cookie не найден", http.StatusUnauthorized)
			return
		}
		accessTokenString := accessTokenCookie.String()

		refreshTokenCookie, err := r.Cookie("refresh_token")
		if err != nil {
			http.Error(w, "refresh_token cookie не найден", http.StatusUnauthorized)
			return
		}
		refreshTokenString := refreshTokenCookie.String()

		if accessTokenString == "" {
			if refreshTokenString == "" {
				http.Error(w, "refresh_token пустой", http.StatusUnauthorized)
				return
			} else if a.refreshTokenPair(r) != nil {
				http.Error(w, "Не удалось обновить пару", http.StatusUnauthorized)
				return
			}
			// Токены обновлены
		}

		valid, err := a.tokenManager.ValidateToken(accessTokenString)
		if !valid {
			http.Error(w, "Некорректный access_token", http.StatusUnauthorized)
			return
		}

		claims, err := a.tokenManager.ParseWithClaims(accessTokenString)
		if err != nil {
			http.Error(w, "Ошибка во время парсинга данных access_token", http.StatusUnauthorized)
			return
		}

		// login := (*claims)["lgn"].(string)
		ip := (*claims)["uip"].(string)

		if r.RemoteAddr != ip {
			http.Error(w, "IP адреса не совпадают", http.StatusUnauthorized)
			return
		}

		if err = a.refreshTokenPair(r); err != nil {
			http.Error(w, "Не удалось обновить пару", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
