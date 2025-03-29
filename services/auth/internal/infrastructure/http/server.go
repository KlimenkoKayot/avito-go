package server

import (
	"fmt"
	"net/http"

	"github.com/klimenkokayot/avito-go/libs/logger"
	"github.com/klimenkokayot/avito-go/libs/router"
	"github.com/klimenkokayot/avito-go/services/auth/config"
	"github.com/klimenkokayot/avito-go/services/auth/internal/domain"
	"github.com/klimenkokayot/avito-go/services/auth/internal/domain/service"
)

// corsMiddleware := cors.New(cors.Options{
// 	AllowedOrigins:   []string{"http://127.0.0.1:8080", "http://localhost:8080"},
// 	AllowedMethods:   []string{"GET", "POST"},
// 	AllowCredentials: true,
// })
// handler := corsMiddleware.Handler(mux)

type AuthServer struct {
	service *service.AuthService
	router  router.Router
	logger  logger.Logger
	cfg     *config.Config
}

func NewAuthServer(service *service.AuthService, cfg *config.Config, logger logger.Logger) (domain.Server, error) {
	router, err := router.NewAdapter(&router.Config{
		Name: cfg.Router,
	})
	if err != nil {
		return nil, err
	}

	server := &AuthServer{
		service,
		router,
		logger,
		cfg,
	}

	err = server.setupRoutes()
	if err != nil {
		return nil, err
	}

	return server, nil
}

func (a *AuthServer) setupRoutes() error {
	return nil
}

func (a *AuthServer) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", a.cfg.ServerPort), a.router)
}
