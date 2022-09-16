package repository

import (
	"github.com/jmoiron/sqlx"
	meme "github.com/mihailco/memessenger"
)

type Authorization interface {
	CreateUser(user meme.User) (int, error)
	GetUser(username, password string) (meme.User, error)
}
type SendMessage interface {
}
type GetMessages interface {
}

type Repository struct {
	Authorization
	SendMessage
	GetMessages
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
