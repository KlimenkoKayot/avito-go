package auth

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	if err := godotenv.Load("internal/auth/.env"); err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
}

func NewDB() (*sqlx.DB, error) {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("Задан пустой DB_DSN для подключения в базе данных")
	}
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDatabaseConn, err.Error())
	}
	return db, nil
}

func CreateUsersTable(db *sqlx.DB) error {
	_, err := db.Exec(`
		DROP TABLE IF EXISTS users;
	`)
	if err != nil {
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
		return err
	}
	log.Println("Table 'users' created successfully")

	return nil
}

func GetPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port, nil
}
