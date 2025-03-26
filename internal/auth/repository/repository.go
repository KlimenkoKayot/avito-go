package auth

import (
	"database/sql"
	"fmt"

	utils "github.com/klimenkokayot/avito-go/internal/auth/utils"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/klimenkokayot/avito-go/pkg/models"
	_ "github.com/lib/pq"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository() (*UserRepository, error) {
	db, err := utils.NewDB()
	if err != nil {
		return nil, err
	}
	return &UserRepository{
		db,
	}, nil
}

func (ur *UserRepository) FindByLogin(login string) (*models.UserSecure, error) {
	user := &models.UserSecure{}
	err := ur.db.Get(user, "SELECT * FROM users WHERE login = $1", login)
	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	} else if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrFindByLogin, err.Error())
	}
	return user, nil
}

func (ur *UserRepository) ExistByLogin(login string) (bool, error) {
	err := ur.db.Get(nil, "SELECT * FROM users WHERE login = $1", login)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("%w: %s", ErrFindByLogin, err.Error())
	}
	return true, nil
}

func (ur *UserRepository) Add(login string, passHash []byte) error {
	found, err := ur.ExistByLogin(login)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrAddUser, err.Error())
	}
	if found {
		return fmt.Errorf("%w: %s", ErrUserExists, login)
	}
	id := uuid.New().String()
	user := &models.UserSecure{
		ID:     id,
		Login:  login,
		Secret: string(passHash),
	}
	_, err = ur.db.NamedExec("INSERT INTO users (id, login, secret) VALUES (:id, :login, :secret)", user)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrAddUser, err.Error())
	}
	return nil
}
