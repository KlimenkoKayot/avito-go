package auth

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := godotenv.Load("internal/view/.env"); err != nil {
		logrus.Fatal("Ошибка загрузки .env файла")
	}
}

func GetPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port, nil
}
