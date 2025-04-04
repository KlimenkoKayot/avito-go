package service

import (
	"github.com/klimenkokayot/avito-go/libs/jwt"
	"github.com/klimenkokayot/avito-go/libs/logger"
	"github.com/klimenkokayot/avito-go/services/auth/config"
	"github.com/klimenkokayot/avito-go/services/auth/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo     domain.UserRepository
	tokenManager *jwt.TokenManager
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
		logger:       logger,
		cfg:          cfg,
	}, nil
}

func (s *AuthService) Register(login, pass string) error {
	secretByte, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	secret := string(secretByte)

	err = s.userRepo.Add(login, secret)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(login, pass string) (string, error) {
	err := s.userRepo.Check(login, pass)
	if err != nil {
		return "", err
	}
	token, err := s.tokenManager.NewRefreshToken(login)
	if err != nil {
		return "", err
	}
	return token, nil
}
