package auth

import (
	"database/sql"
	"fmt"

	utils "github.com/klimenkokayot/avito-go/internal/auth/utils"
	"github.com/sirupsen/logrus"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/klimenkokayot/avito-go/pkg/models"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository() (*UserRepository, error) {
	logrus.Info("Создание UserRepository`я.")
	db, err := utils.NewDB()
	if err != nil {
		logrus.Error("Ошибка при инициализации базы данных.")
		return nil, err
	}
	err = utils.CreateUsersTable(db)
	if err != nil {
		logrus.Error("Ошибка при создании таблицы.")
		return nil, err
	}
	logrus.Debug("Успешно создан UserRepository.")
	return &UserRepository{
		db,
	}, nil
}

func (ur *UserRepository) FindByLogin(login string) (*models.UserSecure, error) {
	logrus.Info("Поиск по логину.")
	user := &models.UserSecure{}
	query := "SELECT * FROM users WHERE login = $1"
	err := ur.db.Get(user, query, login)
	if err == sql.ErrNoRows {
		logrus.Debug("Пользователь не найден.")
		return nil, sql.ErrNoRows
	} else if err != nil {
		logrus.Errorf("Ошибка при запросе в базу данных: %s, login: %s.", query, login)
		return nil, fmt.Errorf("%w: %s", ErrFindByLogin, err.Error())
	}
	logrus.Debug("Успешно.")
	return user, nil
}

func (ur *UserRepository) ExistByLogin(login string) (bool, error) {
	logrus.Info("Проверка на существование по логину.")
	_, err := ur.FindByLogin(login)
	if err == sql.ErrNoRows {
		logrus.Debugf("Успешно: %s не найден.", login)
		return false, nil
	} else if err != nil {
		logrus.Errorf("Ошибка при запросе в базу данных: login: %s.", login)
		return false, err
	}
	logrus.Debugf("Успешно: %s найден.", login)
	return true, nil
}

func (ur *UserRepository) Add(login string, passHash []byte) error {
	logrus.Info("Инициализация нового пользователя.")
	found, err := ur.ExistByLogin(login)
	if err != nil {
		logrus.Error("Ошибка при инициализации нового пользователя.")
		return fmt.Errorf("%w: %s", ErrAddUser, err.Error())
	} else if found {
		logrus.Error("Ошибка, такой пользователь уже существует.")
		return fmt.Errorf("%w: %s", ErrUserExists, login)
	}
	id := uuid.New().String()
	user := &models.UserSecure{
		ID:     id,
		Login:  login,
		Secret: string(passHash),
	}
	logrus.Debug("Запрос на добавление новой строки в таблицу.")
	_, err = ur.db.NamedExec("INSERT INTO users (id, login, secret) VALUES (:id, :login, :secret)", user)
	if err != nil {
		logrus.Error("Ошибка при запросе в таблицу.")
		return fmt.Errorf("%w: %s", ErrAddUser, err.Error())
	}
	logrus.Debugf("Успешно инициализирован: %s.", login)
	return nil
}
