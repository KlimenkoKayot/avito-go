package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenManager struct {
	jwtSecretKey           []byte
	accessTokenExpiration  time.Duration
	refreshTokenExpiration time.Duration
}

func NewTokenManager(jwtSecretKey string, accessTokenExpiration, refreshTokenExpiration time.Duration) (*TokenManager, error) {
	return &TokenManager{
		jwtSecretKey:           []byte(jwtSecretKey),
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

func (tm *TokenManager) ValidateTokenExpiration(token string) (bool, error) {
	valid, err := tm.ValidateToken(token)
	if !valid || err != nil {
		return false, err
	}
	claims, err := tm.ParseWithClaims(token)
	if err != nil {
		return false, err
	}
	expTime := (*claims)["exp"].(time.Time)
	expired := time.Now().After(expTime)
	// если истек, то невалидный
	return !expired, nil
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

/*
Возвращает пару из access (1) и refresh (2) токенов, ошибку (3), если возникла.
*/
func (tm *TokenManager) UpdateTokenPair(refreshToken string, ip string) (string, string, error) {
	valid, err := tm.ValidateToken(refreshToken)
	if err != nil {
		return "", "", err
	}
	if !valid {
		return "", "", ErrNotValidToken
	}

	claims, err := tm.ParseWithClaims(refreshToken)
	if err != nil {
		return "", "", err
	}

	lgn := (*claims)["lgn"].(string)

	refreshToken, err = tm.NewRefreshToken(lgn)
	if err != nil {
		return "", "", err
	}

	accessToken, err := tm.NewAccessToken(lgn, ip)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

/*
Возвращает пару из access (1) и refresh (2) токенов, ошибку (3), если возникла.
*/
func (tm *TokenManager) NewTokenPair(login, ip string) (string, string, error) {
	refreshToken, err := tm.NewRefreshToken(login)
	if err != nil {
		return "", "", err
	}

	accessToken, err := tm.NewAccessToken(login, ip)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
