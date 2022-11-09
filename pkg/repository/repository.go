package repository

import (
	"github.com/jmoiron/sqlx"
	meme "github.com/mihailco/memessenger"
)

type Authorization interface {
	CreateUser(user meme.Userlist) (int, error)
	GetUser(username, password string) (meme.Userlist, error)
	Exist(id int, password string) (bool, error)
}

type Conversation interface {
	Create(creatorId int, info meme.ConversationStruct) (int, error)
	GetAll(userId int) ([]meme.ConversationStruct, error)
	UpdateById(userId int, info meme.ConversationStruct) error
	DeleteById(userId int, itemId int) error
	GetAllUsersAtConv(convId int) ([]meme.Userlist, error)
	IsAMember(userId int, itemId int) bool
	AddUser(userid, convId int) error

	//TODO: add methods
}

type Messages interface {
	Create(info meme.MessageList) (int, error)
}

type Repository struct {
	Authorization
	Conversation
	Messages
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Conversation:  NewConversationPostgres(db),
		Messages:      NewMessagesPostgres(db),
	}
}
