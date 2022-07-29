package meme

import "time"

type Userlist struct {
	Id        int    `json: "id"`
	Password  string `json: "password"`
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
	Phone     string `json: "phone"`
	Email     string `json: "email"`
	Status    string `json: "status"`
	ImageURL  string `json: "imageURL"`
	Username  string `json: "username"`
	IsOnline  bool   `json: "is_online"`
}

type MessageList struct {
	Id            int       `json: "id"`
	Text          string    `json: "text"`
	UserIdFrom    string    `json: "user_id_from"`
	UserIdTo      string    `json: "user_id_to"`
	CreatedAt     time.Time `json: "created_at"`
	TypeOfMessage string    `json: "type_of_message"`
}

type GroupMemember struct {
	GroupId string `json: "group_id"`
	UserId  string `json: "user_id"`
}
