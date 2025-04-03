package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenManager struct {
	jwtSecretKey           string
	accessTokenExpiration  time.Duration
	refreshTokenExpiration time.Duration
}

func NewTokenManager(jwtSecretKey string, accessTokenExpiration, refreshTokenExpiration time.Duration) (*TokenManager, error) {
	return &TokenManager{
		jwtSecretKey:           jwtSecretKey,
		accessTokenExpiration:  accessTokenExpiration,
		refreshTokenExpiration: refreshTokenExpiration,
	}, nil
}

func (tm *TokenManager) NewAccessToken(login string, ip string) (string, error) {
	payload := jwt.MapClaims{
		"lgn": login,
		"uip": ip,
		"exp": time.Now().Add(tm.accessTokenExpiration).Unix(),
		"ctd": time.Now().Unix(),
	}
	tokenData, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString(tm.jwtSecretKey)
	if err != nil {
		return "", err
	}
	return tokenData, nil
}

func (tm *TokenManager) NewRefreshToken(login string) (string, error) {
	payload := jwt.MapClaims{
		"lgn": login,
		"exp": time.Now().Add(tm.refreshTokenExpiration).Unix(),
		"ctd": time.Now().Unix(),
	}
	tokenData, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString(tm.jwtSecretKey)
	if err != nil {
		return "", err
	}
	return tokenData, nil
}

func (tm *TokenManager) ValidateToken(tokenString string) (bool, error) {
	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return tm.jwtSecretKey, nil
	})
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}

func (tm *TokenManager) ParseWithClaims(tokenString string) (*jwt.MapClaims, error) {
	claims := &jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return tm.jwtSecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
