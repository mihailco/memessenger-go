package meme

import (
	"time"
)

type Message struct {
	Id            int       `json: "id"`
	Text          string    `json: "text"`
	UserIdFrom    string    `json: "userIdFrom"`
	UserIdTo      string    `json: "userIdTo"`
	CreatedAt     time.Time `json: "createdAt"`
	TypeOfMessage string    `json: "typeOfMessage"`
}
