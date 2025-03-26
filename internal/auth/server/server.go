package auth

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/klimenkokayot/avito-go/internal/auth/server/handlers"
	utils "github.com/klimenkokayot/avito-go/internal/auth/utils"
)

type AuthServer struct {
	handler *handlers.AuthHandler
}

func NewAuthServer() (*AuthServer, error) {
	handler, err := handlers.NewAuthHandler()
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrNewServer, err.Error())
	}
	return &AuthServer{
		handler,
	}, nil
}

func (s *AuthServer) Run() error {
	port, err := utils.GetPort()
	if err != nil {
		return fmt.Errorf("%w: %s", ErrRunServer, err.Error())
	}

	mux := mux.NewRouter()
	mux.HandleFunc("/auth/register", s.handler.Register).Methods("POST")
	return http.ListenAndServe(":"+port, mux)
}
