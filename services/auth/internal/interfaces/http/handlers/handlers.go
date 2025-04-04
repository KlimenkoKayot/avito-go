package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/klimenkokayot/avito-go/libs/logger"
	"github.com/klimenkokayot/avito-go/services/auth/config"
	"github.com/klimenkokayot/avito-go/services/auth/internal/domain"
	"github.com/klimenkokayot/avito-go/services/auth/internal/domain/service"
)

type AuthHandler struct {
	authService *service.AuthService
	logger      logger.Logger
	cfg         *config.Config
}

func NewAuthHandler(service *service.AuthService, cfg *config.Config, logger logger.Logger) (*AuthHandler, error) {
	logger.Info("Инициализация обработчика.")
	logger.OK("Обработчик успешно инициализирован.")
	return &AuthHandler{
		service,
		logger,
		cfg,
	}, nil
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error("Ошибка при чтении тела запроса.", logger.Field{
			Key:   "err",
			Value: err.Error(),
		})
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Errorf("ошибка при чтении тела запроса: %w", err))
		return
	}
	defer r.Body.Close()

	user := &domain.User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		h.logger.Error("Ошибка при парсинге тела запроса.", logger.Field{
			Key:   "err",
			Value: err.Error(),
		})
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(fmt.Errorf("ошибка при парсинге тела запроса: %w", err))
		return
	}

	err = h.authService.Register(user.Login, user.Secret)
	if err != nil {
		h.logger.Error("Ошибка при регистрации пользователя", logger.Field{
			Key:   "err",
			Value: err.Error(),
		})
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(fmt.Errorf("ошибка при регистрации: %w", err))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.Writer(w).Write([]byte(fmt.Errorf("ошибка при создании AuthHandler: %w", err).Error()))
		return
	}
	defer r.Body.Close()

	user := &domain.User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		h.logger.Error("Ошибка при парсинге тела запроса.", logger.Field{
			Key:   "err",
			Value: err.Error(),
		})
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(fmt.Errorf("ошибка при парсинге тела запроса: %w", err))
		return
	}

	token, err := h.authService.Login(user.Login, user.Secret)
	if err != nil {
		h.logger.Error("Неудачный вход в аккаунт пользователя", logger.Field{
			Key:   "err",
			Value: err.Error(),
		})
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(fmt.Errorf("ошибка при попытке входа: %w", err))
		return
	}

	r.AddCookie(&http.Cookie{
		Name:  "refresh_token",
		Value: token,
	})

	w.WriteHeader(http.StatusOK)
}
