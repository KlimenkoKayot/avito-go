package models

import "time"

type User struct {
	Login    string `json:"login"`
	Password string `json:"pass"`
}

type UserSecure struct {
	ID        string    `json:"id" db:"id"`
	Login     string    `json:"login" db:"login"`
	Secret    string    `json:"secret" db:"secret"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
