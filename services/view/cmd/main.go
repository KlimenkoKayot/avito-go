package main

import (
	server "github.com/klimenkokayot/avito-go/services/view/internal/server"
	utils "github.com/klimenkokayot/avito-go/services/view/pkg/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Запуск микросервиса View.")
	server, err := server.NewViewServer()
	if err != nil {
		logrus.Fatal("Неудачная инициализация ViewServer`a.")
	}
	port, err := utils.GetPort()
	if err != nil {
		logrus.Fatal("Ошибка при получении порта.")
	}
	logrus.Debugf("Сервер запущен на порту %s\n.", port)
	if err := server.Run(); err != nil {
		logrus.Fatal("Ошибка при запуске сервера.")
	}
}
