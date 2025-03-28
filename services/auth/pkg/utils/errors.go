package auth

import "fmt"

var (
	ErrDatabaseConn = fmt.Errorf("ошибка при подключении к базе данных (DB_DSN)")
)
