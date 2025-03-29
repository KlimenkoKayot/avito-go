package repo

import "fmt"

var (
	ErrUserExists = fmt.Errorf("пользователь с таким логином существует")
	ErrBadDSN     = fmt.Errorf("неправильный database dsn")
)
