package repository

type Authorization interface {
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

func NewRepository() *Repository {
	return &Repository{}
}
