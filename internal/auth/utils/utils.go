package auth

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := godotenv.Load("internal/auth/.env"); err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
}

func NewDB() (*sqlx.DB, error) {
	logrus.Info("Инициализация базы данных.")
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		logrus.Error("Задан пустой DB_DSN для подключения в базе данных")
		return nil, ErrDatabaseConn
	}
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logrus.Errorf("Ошибка при подключении по DSN: %s.", err.Error())
		return nil, fmt.Errorf("%w: %s", ErrDatabaseConn, err.Error())
	}
	logrus.Debug("Инициализация успешна.")
	return db, nil
}

func CreateUsersTable(db *sqlx.DB) error {
	logrus.Info("Инициализация таблицы users.")
	_, err := db.Exec(`
		DROP TABLE IF EXISTS users;
	`)
	if err != nil {
		logrus.Error("Ошибка при запросе в базу данных.")
		return err
	}

	_, err = db.Exec(`
			CREATE TABLE users (
				id UUID PRIMARY KEY,
				login TEXT UNIQUE NOT NULL,
				secret TEXT NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			)
		`)
	if err != nil {
		logrus.Error("Ошибка при запросе в базу данных.")
		return err
	}

	logrus.Debug("Успешно создана таблица users.")
	return nil
}

func GetPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port, nil
}
