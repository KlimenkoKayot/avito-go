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
	logger.OK("Успешно.")
	return &AuthHandler{
		service,
		logger,
		cfg,
	}, nil
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.Writer(w).Write([]byte(fmt.Errorf("%w: %s", ErrReadBody, err.Error()).Error()))
		return
	}
	defer r.Body.Close()

	user := &domain.User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		io.Writer(w).Write([]byte(fmt.Errorf("%w: %s", ErrUnprocessibleEntity, err.Error()).Error()))
		return
	}

	err = h.authService.Register(user.Login, user.Secret)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		io.Writer(w).Write([]byte(fmt.Errorf("%w: %s", ErrRegisterProblem, err.Error()).Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.Writer(w).Write([]byte(fmt.Errorf("%w: %s", ErrReadBody, err.Error()).Error()))
		return
	}
	defer r.Body.Close()

	user := &domain.User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		io.Writer(w).Write([]byte(fmt.Errorf("%w: %s", ErrUnprocessibleEntity, err.Error()).Error()))
		return
	}

	err = h.authService.Login(user.Login, user.Secret)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		io.Writer(w).Write([]byte(fmt.Errorf("%w: %s", ErrRegisterProblem, err.Error()).Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
