package repo

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/klimenkokayot/avito-go/libs/logger"
	"github.com/klimenkokayot/avito-go/services/auth/config"
	"github.com/klimenkokayot/avito-go/services/auth/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	logger logger.Logger
	db     *sqlx.DB
}

func NewUserRepository(cfg *config.Config, logger logger.Logger) (domain.UserRepository, error) {
	logger.Info("Инициализация user-репозитория.")
	dsn := cfg.DatabaseDSN
	if dsn == "" {
		return nil, fmt.Errorf("лох!!... пидр!...")
	}

	logger.Info("Подключение по DSN.")
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`DROP TABLE IF EXISTS users;`)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	logger.OK("Успешно.")
	return &UserRepository{
		logger,
		db,
	}, nil
}

func (ur *UserRepository) FindByLogin(login string) (*domain.User, error) {
	user := &domain.User{}
	query := "SELECT * FROM users WHERE login = $1"
	err := ur.db.Get(user, query, login)
	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	} else if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) ExistByLogin(login string) (bool, error) {
	_, err := ur.FindByLogin(login)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func (ur *UserRepository) Add(login string, secret string) error {
	found, err := ur.ExistByLogin(login)
	if err != nil {
		return err
	} else if found {
		return fmt.Errorf("АШИПКААААААААА_сосал?")
	}
	id := uuid.New().String()
	user := &domain.User{
		ID:     id,
		Login:  login,
		Secret: secret,
	}
	_, err = ur.db.NamedExec("INSERT INTO users (id, login, secret) VALUES (:id, :login, :secret)", user)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) Check(login, pass string) error {
	var secret string
	err := ur.db.QueryRow("SELECT secret FROM users WHERE login = $1", login).Scan(&secret)
	if err == sql.ErrNoRows {
		return sql.ErrNoRows
	} else if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(secret), []byte(pass))
	if err != nil {
		return err
	}
	return nil
}
