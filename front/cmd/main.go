package main

import (
	"log"

	"github.com/klimenkokayot/avito-go/front/config"
	"github.com/klimenkokayot/avito-go/front/internal/app"
	"github.com/klimenkokayot/avito-go/libs/logger"
)

func main() {
	cfg, err := config.Load("")
	if err != nil {
		log.Fatalf("Ошибка при инициализации конфига: %s", err.Error())
	}

	logger, err := logger.NewAdapter(&logger.Config{
		Adapter: cfg.Logger,
		Level:   logger.LevelDebug,
	})

	appLogger := logger.WithLayer("APP")

	app, err := app.NewApplication(cfg, appLogger)
	if err != nil {
		appLogger.Fatal(err.Error())
	}

	if err := app.Run(); err != nil {
		appLogger.Fatal(err.Error())
	}
}
