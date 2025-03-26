package auth

import (
	"fmt"

	repo "github.com/klimenkokayot/avito-go/internal/auth/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repo.UserRepository
}

func NewAuthService() (*AuthService, error) {
	repo, err := repo.NewUserRepository()
	if err != nil {
		return nil, err
	}
	return &AuthService{
		repo,
	}, nil
}

func (s *AuthService) Register(login, pass string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrGenerateFromPass, err.Error())
	}

	err = s.userRepo.Add(login, hash)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrAddNewUser, err.Error())
	}

	return nil
}
