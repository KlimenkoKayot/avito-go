package main

import (
	"github.com/joho/godotenv"
	server "github.com/klimenkokayot/avito-go/front/internal/server"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("Ошибка загрузки .env файла: %s.", err.Error())
	}
}

func main() {
	logrus.Info("Запуск микросервиса View.")
	server, err := server.NewViewServer()
	if err != nil {
		logrus.Fatal("Неудачная инициализация ViewServer`a.")
	}
	logrus.Debugf("Сервер запущен на порту %s\n.", port)
	if err := server.Run(); err != nil {
		logrus.Fatal("Ошибка при запуске сервера.")
	}
}
