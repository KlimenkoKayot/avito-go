package main

import (
	"log"

	"github.com/klimenkokayot/avito-go/services/auth/config"
	"github.com/klimenkokayot/avito-go/services/auth/internal/app"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Ошибка при инициализации логгера: %s.", err.Error())
		return
	}

	config, err := config.NewConfig()
	if err != nil {
		logger.Sugar().Fatalf("Ошибка при инициализации конфига: %s.", err.Error())
	}

	app, err := app.NewApplication(config, logger)
}
