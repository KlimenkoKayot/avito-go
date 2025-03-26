package auth

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("internal/view/.env"); err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
}

func GetPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port, nil
}
