package auth

import "fmt"

var (
	ErrNewServer = fmt.Errorf("ошибка при создании сервера")
	ErrRunServer = fmt.Errorf("ошибка при запуске сервера")
)
