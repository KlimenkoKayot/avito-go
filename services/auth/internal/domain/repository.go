package domain

import "time"

type User struct {
	ID        string
	Login     string
	Secret    string
	CreatedAt time.Time
}

type UserRepository interface {
	Add(login string, secret string) error
	Check(login string, pass string) error
	FindByLogin(login string) (*User, error)
	ExistByLogin(login string) (bool, error)
}
