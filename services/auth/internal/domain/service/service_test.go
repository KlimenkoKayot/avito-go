package service_test

import (
	"testing"

	"github.com/klimenkokayot/avito-go/services/auth/config"
	"github.com/klimenkokayot/avito-go/services/auth/internal/domain/service"
	"github.com/klimenkokayot/avito-go/services/auth/mocks/jwt"
	"github.com/klimenkokayot/avito-go/services/auth/mocks/logger"
	"github.com/klimenkokayot/avito-go/services/auth/mocks/repository"
	"go.uber.org/mock/gomock"
)

func TestAuthService_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	mockTokenManager := jwt.NewMockTokenManager(ctrl)
	mockLogger := logger.NewMockLogger(ctrl)

	cfg := &config.Config{
		JwtSecretKey:           "secret",
		AccessTokenExpiration:  3600,
		RefreshTokenExpiration: 7200,
	}

	// Устанавливаем ожидание для метода Info
	mockLogger.EXPECT().Info("Инициализация сервиса.").Times(1)
	mockLogger.EXPECT().OK("Успешно.").Times(1)

	authService, err := service.NewAuthService(mockRepo, mockTokenManager, cfg, mockLogger)
	if err != nil {
		t.Fatalf("Failed to create AuthService: %v", err)
	}

	login := "testuser"
	password := "password"

	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// Изменяем ожидание на использование gomock.Any() для второго аргумента
	mockRepo.EXPECT().Add(login, gomock.Any()).Return(nil)
	mockTokenManager.EXPECT().NewTokenPair(gomock.Any(), gomock.Any()).Return("access_token", "refresh_token", nil)

	accessToken, refreshToken, err := authService.Register(login, password)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if accessToken == "" || refreshToken == "" {
		t.Errorf("Tokens should not be empty")
	}
}
