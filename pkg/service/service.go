package service

import "github.com/mihailco/memessenger/pkg/repository"

type Authorization interface {
}
type SendMessage interface {
}
type GetMessages interface {
}

type Service struct {
	Authorization
	SendMessage
	GetMessages
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
