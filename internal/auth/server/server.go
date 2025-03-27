package auth

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/klimenkokayot/avito-go/internal/auth/server/handlers"
	utils "github.com/klimenkokayot/avito-go/internal/auth/utils"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

type AuthServer struct {
	handler *handlers.AuthHandler
}

func NewAuthServer() (*AuthServer, error) {
	logrus.Debug("Инициализация AuthServer`a.")
	handler, err := handlers.NewAuthHandler()
	if err != nil {
		logrus.Errorf("Ошибка при создании AuthServer`а.")
		return nil, fmt.Errorf("%w: %s", ErrNewServer, err.Error())
	}
	logrus.Debug("Успешно создан AuthServer.")
	return &AuthServer{
		handler,
	}, nil
}

func (s *AuthServer) Run() error {
	logrus.Info("Запуск AuthServer`a сервера.")
	port, err := utils.GetPort()
	if err != nil {
		return fmt.Errorf("%w: %s", ErrRunServer, err.Error())
	}
	logrus.Debugf("Получен порт: %s.", port)

	mux := mux.NewRouter()
	mux.HandleFunc("/auth/register", s.handler.Register).Methods("POST")
	mux.HandleFunc("/auth/login", s.handler.Login).Methods("POST")

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:8080", "http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
	})
	handler := corsMiddleware.Handler(mux)
	logrus.Debug("Cors настроен.")

	return http.ListenAndServe(":"+port, handler)
}
