package view

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	handlers "github.com/klimenkokayot/avito-go/internal/view/server/handlers"
	utils "github.com/klimenkokayot/avito-go/internal/view/utils"
)

type ViewServer struct {
	H *handlers.ViewHandler
}

func NewViewServer() (*ViewServer, error) {
	handlers, err := handlers.NewAuthHandler()
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrNewServer, err.Error())
	}
	return &ViewServer{
		handlers,
	}, nil
}

func (s *ViewServer) Run() error {
	port, err := utils.GetPort()
	if err != nil {
		return fmt.Errorf("%w: %s", ErrRunServer, err.Error())
	}

	mux := mux.NewRouter()
	mux.HandleFunc("/login", s.H.LoginPage).Methods("GET")
	mux.HandleFunc("/register", s.H.RegisterPage).Methods("GET")

	staticDir := filepath.Join("internal", "view", "web", "static")
	fs := http.FileServer(http.Dir(staticDir))
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	return http.ListenAndServe(":"+port, mux)
}
