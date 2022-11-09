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

func (s *ConversationService) UpdateById(userId int, info meme.ConversationStruct) error {
	return s.repo.UpdateById(userId, info)
}

func (s *ConversationService) DeleteById(userId int, itemId int) error {
	return s.repo.DeleteById(userId, itemId)
}

func (s *ConversationService) GetAllUsersAtConv(convId int) ([]meme.Userlist, error) {
	return s.repo.GetAllUsersAtConv(convId)
}

func (s *ConversationService) IsAMember(userId int, itemId int) bool {
	return s.repo.IsAMember(userId, itemId)
}

func (s *ConversationService) AddUser(userid, convId int) error {
	return s.repo.AddUser(userid, convId)
}
