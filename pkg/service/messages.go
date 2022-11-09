package service

import (
	meme "github.com/mihailco/memessenger"
	"github.com/mihailco/memessenger/pkg/repository"
)

type MessageService struct {
	repo repository.Messages
}

func NewMessageService(repo repository.Messages) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) Create(info meme.MessageList) (int, error) {
	return s.repo.Create(info)
}
