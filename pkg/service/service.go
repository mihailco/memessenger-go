package service

import (
	meme "github.com/mihailco/memessenger"
	"github.com/mihailco/memessenger/pkg/repository"
)

type Authorization interface {
	CreateUser(user meme.User) (int, error)
	GenerateToken(username, password string) (string, error)
}
type SendMessage interface {
}
type GetMessages interface {
}

type Service struct {
	Authorization
	// SendMessage
	// GetMessages
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
