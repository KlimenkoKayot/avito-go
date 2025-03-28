package auth

import (
	"fmt"

	repo "github.com/klimenkokayot/avito-go/internal/auth/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repo.UserRepository
}

func NewAuthService() (*AuthService, error) {
	logrus.Info("Инициализация AuthService`a.")
	repo, err := repo.NewUserRepository()
	if err != nil {
		logrus.Error("Ошибка при создании AuthService`a!")
		return nil, err
	}
	logrus.Debug("Успешно создан AuthService.")
	return &AuthService{
		repo,
	}, nil
}

func (s *AuthService) Register(login, pass string) error {
	logrus.Info("Запрос на регистрацию в сервис.")
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("Ошибка при генерации секрета: %s", pass)
		return fmt.Errorf("%w: %s", ErrGenerateFromPass, err.Error())
	}

	err = s.userRepo.Add(login, hash)
	if err != nil {
		logrus.Error("Ошибка при инициализации пользователя.")
		return fmt.Errorf("%w: %s", ErrAddNewUser, err.Error())
	}

	logrus.Debug("Сервис успешно зарегистрировал пользователя.")
	return nil
}

func (s *AuthService) Login(login, pass string) error {
	logrus.Info("Запрос на вход в сервис.")
	err := s.userRepo.Check(login, pass)
	if err != nil {
		logrus.Error("Ошибка при инициализации пользователя.")
		return fmt.Errorf("%w: %s", ErrAddNewUser, err.Error())
	}

	logrus.Debug("Сервис успешно обработал вход пользователя.")
	return nil
}
