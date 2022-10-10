package repository

import (
	"github.com/jmoiron/sqlx"
	meme "github.com/mihailco/memessenger"
)

type Authorization interface {
	CreateUser(user meme.User) (int, error)
	GetUser(username, password string) (meme.User, error)
}

type Conversation interface {
	Create(creatorId int, info meme.ConversationStruct) (int, error)
	GetAll(userId int) ([]meme.ConversationStruct, error)

	//TODO: add methods
}

type SendMessage interface {
}
type GetMessages interface {
}

type Repository struct {
	Authorization
	Conversation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Conversation:  NewConversationListPostgres(db),
	}
}
