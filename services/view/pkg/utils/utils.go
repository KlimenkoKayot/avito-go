package view

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("Ошибка загрузки .env файла: %s.", err.Error())
	}
}

func GetPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port, nil
}
