package service

import (
	meme "github.com/mihailco/memessenger"
	"github.com/mihailco/memessenger/pkg/repository"
)

type Authorization interface {
	CreateUser(user meme.Userlist) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, string, error)
	Exist(id int, password string) (bool, error)
}

type Conversation interface {
	Create(creatorId int, info meme.ConversationStruct) (int, error)
	GetAll(id int) ([]meme.ConversationStruct, error)
	UpdateById(userId int, indo meme.ConversationStruct) error
	DeleteById(userId int, itemId int) error
	GetAllUsersAtConv(convId int) ([]meme.Userlist, error)
	IsAMember(userId int, itemId int) bool
	AddUser(userid, convId int) error
	//TODO: get users from conversation
}

type Messages interface {
	Create(info meme.MessageList) (int, error)
}

type Service struct {
	Authorization
	Conversation
	Messages
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Conversation:  NewConversationService(repos.Conversation),
		Messages:      NewMessageService(repos.Messages),
	}
}
