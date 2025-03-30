package service

import (
	"github.com/klimenkokayot/avito-go/libs/logger"
	"github.com/klimenkokayot/avito-go/services/auth/config"
	"github.com/klimenkokayot/avito-go/services/auth/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	logger   logger.Logger
	userRepo domain.UserRepository
}

func NewAuthService(repo domain.UserRepository, cfg *config.Config, logger logger.Logger) (*AuthService, error) {
	logger.Info("Инициализация сервиса.")
	logger.OK("Успешно.")
	return &AuthService{
		logger,
		repo,
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

func (s *AuthService) Login(login, pass string) error {
	err := s.userRepo.Check(login, pass)
	if err != nil {
		return err
	}
	return nil
}
