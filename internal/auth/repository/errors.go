package auth

import "fmt"

var (
	ErrFindByLogin = fmt.Errorf("ошибка при поиске пользователя по логину")
	ErrAddUser     = fmt.Errorf("ошибка при добавлении нового пользователя")
	ErrUserExists  = fmt.Errorf("пользователь с таким логином существует")
)
