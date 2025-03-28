package handlers

import "fmt"

var (
	ErrCreateAuthHandler   = fmt.Errorf("ошибка при создании AuthHandler")
	ErrReadBody            = fmt.Errorf("ошибка при чтении тела запроса")
	ErrUnprocessibleEntity = fmt.Errorf("ошибка при парсинге тела запроса")
	ErrRegisterProblem     = fmt.Errorf("ошибка при регистрации нового пользователя")
)
