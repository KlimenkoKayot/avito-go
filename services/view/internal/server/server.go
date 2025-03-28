package view

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	handlers "github.com/klimenkokayot/avito-go/services/view/internal/server/handlers"
	utils "github.com/klimenkokayot/avito-go/services/view/pkg/utils"
	"github.com/sirupsen/logrus"
)

type ViewServer struct {
	H *handlers.ViewHandler
}

func NewViewServer() (*ViewServer, error) {
	logrus.Debug("Инициализация ViewServer`a.")
	handlers, err := handlers.NewViewHandler()
	if err != nil {
		logrus.Errorf("Ошибка при создании ViewServer`a.")
		return nil, fmt.Errorf("%w: %s", ErrNewServer, err.Error())
	}
	logrus.Debug("Успешно создан ViewServer.")
	return &ViewServer{
		handlers,
	}, nil
}

func (s *ViewServer) Run() error {
	logrus.Info("Запуск ViewServer`a сервера.")
	port, err := utils.GetPort()
	if err != nil {
		return fmt.Errorf("%w: %s", ErrRunServer, err.Error())
	}
	logrus.Debugf("Получен порт: %s.", port)

	mux := mux.NewRouter()
	mux.HandleFunc("/login", s.H.LoginPage).Methods("GET")
	mux.HandleFunc("/register", s.H.RegisterPage).Methods("GET")

	staticDir := filepath.Join("internal", "view", "web", "static")
	fs := http.FileServer(http.Dir(staticDir))
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	return http.ListenAndServe(":"+port, mux)
}
