package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/klimenkokayot/avito-go/libs/logger"
	"github.com/klimenkokayot/avito-go/services/auth/config"
)

type TokenManager struct {
	logger logger.Logger
	cfg    *config.Config
}

func NewTokenManager(jwtSecretKey string, logger logger.Logger, cfg *config.Config) (*TokenManager, error) {
	return &TokenManager{
		logger: logger,
		cfg:    cfg,
	}, nil
}

func (tm *TokenManager) NewAccessToken(login string, ip string) (string, error) {
	payload := jwt.MapClaims{
		"lgn": login,
		"uip": ip,
		"exp": time.Now().Add(tm.cfg.AccessTokenExpiration).Unix(),
		"ctd": time.Now().Unix(),
	}
	tokenData, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString(tm.cfg.JwtSecretKey)
	if err != nil {
		return "", err
	}
	return tokenData, nil
}

func (tm *TokenManager) NewRefreshToken(login string) (string, error) {
	payload := jwt.MapClaims{
		"lgn": login,
		"exp": time.Now().Add(tm.cfg.RefreshTokenExpiration).Unix(),
		"ctd": time.Now().Unix(),
	}
	tokenData, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString(tm.cfg.JwtSecretKey)
	if err != nil {
		return "", err
	}
	return tokenData, nil
}

func (tm *TokenManager) ValidateToken(tokenString string) (bool, error) {
	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return tm.cfg.JwtSecretKey, nil
	})
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}

func (tm *TokenManager) ParseWithClaims(tokenString string) (*jwt.MapClaims, error) {
	claims := &jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return tm.cfg.JwtSecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
