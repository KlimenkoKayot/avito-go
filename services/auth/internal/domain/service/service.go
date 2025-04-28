package service

import (
	"context"
	"sync"

	"github.com/klimenkokayot/avito-go/libs/jwt"
	"github.com/klimenkokayot/avito-go/libs/logger"
	"github.com/klimenkokayot/avito-go/services/auth/config"
	"github.com/klimenkokayot/avito-go/services/auth/internal/domain/model"
	domain "github.com/klimenkokayot/avito-go/services/auth/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo     domain.UserRepository
	tokenManager *jwt.TokenManager
	mu           sync.Mutex
	logger       logger.Logger
	cfg          *config.Config
}

func NewAuthService(repo domain.UserRepository, cfg *config.Config, logger logger.Logger) (*AuthService, error) {
	logger.Info("Инициализация сервиса.")
	tokenManager, err := jwt.NewTokenManager(cfg.JwtSecretKey, cfg.AccessTokenExpiration, cfg.RefreshTokenExpiration)
	if err != nil {
		return nil, err
	}
	logger.OK("Успешно.")
	return &AuthService{
		userRepo:     repo,
		tokenManager: tokenManager,
		mu:           sync.Mutex{},
		logger:       logger,
		cfg:          cfg,
	}, nil
}

func (s *AuthService) Register(login, pass string) (string, string, error) {
	secretByte, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}
	secret := string(secretByte)

	err = s.userRepo.Add(login, secret)
	if err != nil {
		return "", "", err
	}

	accessToken, refreshToken, err := s.tokenManager.NewTokenPair(login)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) Login(login, pass string) (string, string, error) {
	valid, err := s.userRepo.Check(login, pass)
	if !valid || err != nil {
		return "", "", domain.ErrBadPassword
	}

	accessToken, refreshToken, err := s.tokenManager.NewTokenPair(login)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) ValidateTokenPair(ctx context.Context, tokenPair *model.TokenPair) (bool, error) {
	validAccessToken, err := s.tokenManager.ValidateTokenExpiration(tokenPair.AccessToken)
	if err != nil {
		return false, err
	}
	validRefreshToken, err := s.tokenManager.ValidateTokenExpiration(tokenPair.RefreshToken)
	if err != nil {
		return false, err
	}
	validPairToken := validAccessToken && validRefreshToken
	return validPairToken, nil
}
