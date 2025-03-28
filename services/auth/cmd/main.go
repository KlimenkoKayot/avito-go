package main

import (
	server "github.com/klimenkokayot/avito-go/services/auth/internal/server"
	utils "github.com/klimenkokayot/avito-go/services/auth/pkg/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Запуск микросервиса Auth.")
	server, err := server.NewAuthServer()
	if err != nil {
		logrus.Fatal("Неудачная инициализация AuthServer`a.")
	}
	logrus.Debug("Запрос на получение порта из .env.")
	port, err := utils.GetPort()
	if err != nil {
		logrus.Fatal("Ошибка при получении порта.")
	}
	logrus.Debugf("Сервер запущен на порту %s.", port)
	if err := server.Run(); err != nil {
		logrus.Fatal("Ошибка при запуске сервера.")
	}
}
