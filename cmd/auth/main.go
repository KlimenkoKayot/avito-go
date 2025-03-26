package main

import (
	"log"

	server "github.com/klimenkokayot/avito-go/internal/auth/server"
	utils "github.com/klimenkokayot/avito-go/internal/auth/utils"
)

func main() {
	server, err := server.NewAuthServer()
	if err != nil {
		log.Fatalf("Ошибка: %s\n", err.Error())
	}
	port, _ := utils.GetPort()
	log.Printf("Сервер запущен на порту %s\n", port)
	if err := server.Run(); err != nil {
		log.Fatalf("Ошибка сервера: %s\n", err.Error())
	}
}
