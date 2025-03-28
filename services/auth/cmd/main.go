package main

import (
	"log"

	"github.com/klimenkokayot/avito-go/libs/logger"
	"github.com/klimenkokayot/avito-go/services/auth/config"
	"github.com/klimenkokayot/avito-go/services/auth/internal/app"
)

func main() {
	config, err := config.Load("")
	if err != nil {
		log.Fatalf("Ошибка при инициализации config`a: %s.", err.Error())
	}

	logger, err := logger.NewAdapter()

	if err != nil {
		log.Fatalf("Ошибка при инициализации config`a: %s.", err.Error())
	}

	app, err := app.NewApplication(config, logger)
}
