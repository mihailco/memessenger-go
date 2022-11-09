package meme

import "time"

type Userlist struct {
	UserId    int    `json:"id" db:"user_id"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	ImageURL  string `json:"imageURL"`
	Username  string `json:"username"`
	IsOnline  bool   `json:"is_online"`
}

type MessageList struct {
	Id               int       `json:"id" db:"id"`
	Text             string    `json:"text"`
	UserIdFrom       int       `json:"userId_from"`
	CreatedAt        time.Time `json:"created_at"`
	TypeOfMessage    string    `json:"type_of_message"`
	ConversationIdTo int       `json:"conversation_id_to"`
}

type GroupMemember struct {
	GroupId string `json:"group_id"`
	UserId  string `json:"user_id"`
}

type ConversationStruct struct {
	Id          int       `json:"id" db:"id"`
	ChannelName string    `json:"channel_name" db:"channel_name"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Description string    `json:"channel_description" db:"channel_description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	CreatorId   int       `json:"creator_id" db:"creator_id"`
}
