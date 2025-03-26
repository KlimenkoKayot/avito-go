package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	service "github.com/klimenkokayot/avito-go/internal/auth/service"
	"github.com/klimenkokayot/avito-go/pkg/models"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler() (*AuthHandler, error) {
	authService, err := service.NewAuthService()
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrCreateAuthHandler, err.Error())
	}
	return &AuthHandler{
		authService,
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

	user := &models.User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		io.Writer(w).Write([]byte(fmt.Errorf("%w: %s", ErrUnprocessibleEntity, err.Error()).Error()))
		return
	}

	err = h.authService.Register(user.Login, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		io.Writer(w).Write([]byte(fmt.Errorf("%w: %s", ErrRegisterProblem, err.Error()).Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
