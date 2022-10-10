package service

import (
	meme "github.com/mihailco/memessenger"
	"github.com/mihailco/memessenger/pkg/repository"
)

type ConversationService struct {
	repo repository.Conversation
}

func NewConversationService(repo repository.Conversation) *ConversationService {
	return &ConversationService{repo: repo}
}

func (s *ConversationService) Create(creatorId int, info meme.ConversationStruct) (int, error) {
	return s.repo.Create(creatorId, info)
}

func (s *ConversationService) GetAll(id int) ([]meme.ConversationStruct, error) {
	return s.repo.GetAll(id)
}
