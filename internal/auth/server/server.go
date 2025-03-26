package auth

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/klimenkokayot/avito-go/internal/auth/server/handlers"
	utils "github.com/klimenkokayot/avito-go/internal/auth/utils"
	"github.com/rs/cors"
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

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:8080", "http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
	})
	handler := corsMiddleware.Handler(mux)

	return http.ListenAndServe(":"+port, handler)
}
