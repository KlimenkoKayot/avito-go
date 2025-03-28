package auth

import "fmt"

var (
	ErrGenerateFromPass = fmt.Errorf("ошибка при генерации пароля bcrypt")
	ErrAddNewUser       = fmt.Errorf("ошибка при создании нового пользователя")
)
