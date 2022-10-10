package service

import (
	meme "github.com/mihailco/memessenger"
	"github.com/mihailco/memessenger/pkg/repository"
)

type Authorization interface {
	CreateUser(user meme.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Conversation interface {
	Create(creatorId int, info meme.ConversationStruct) (int, error)
	GetAll(id int) ([]meme.ConversationStruct, error)
}

type Service struct {
	Authorization
	Conversation
	// SendMessage
	// GetMessages
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Conversation:  NewConversationService(repos),
	}
}
