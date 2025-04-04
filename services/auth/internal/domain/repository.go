package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Login     string    `db:"login" json:"login"`
	Secret    string    `db:"secret" json:"pass"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type UserRepository interface {
	Add(login string, secret string) error
	Check(login string, pass string) error
	FindByLogin(login string) (*User, error)
	ExistByLogin(login string) (bool, error)
}
