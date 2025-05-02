package main

import (
	"log"

	"github.com/klimenkokayot/avito-go/api_gateway/config"
	"github.com/klimenkokayot/avito-go/api_gateway/internal/app"
	"github.com/klimenkokayot/avito-go/libs/logger"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	logger, err := logger.NewAdapter(&logger.Config{
		Adapter: cfg.Router,
		Level:   logger.LevelInfo,
	})
	if err != nil {
		log.Fatal(err)
	}

	application, err := app.NewApplication(logger, cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
