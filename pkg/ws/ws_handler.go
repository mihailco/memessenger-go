package ws

import (
	"encoding/json"
	"fmt"

	meme "github.com/mihailco/memessenger"
	"github.com/sirupsen/logrus"
)

type EventType struct {
	Mestype   string `json:"opcode"`
	AuthToken string `json:"token"`
}

func (h *Hub) WSHandler(message []byte) error {
	var event EventType

	if err := json.Unmarshal(message, &event); err != nil {
		logrus.Println("error in Unmarshal")
	}
	switch event.Mestype {
	case "sendmessage":
		var mes meme.MessageList

		err := json.Unmarshal(message, &mes)
		if err != nil {
			fmt.Println(err)
			return err
		}

		userId, _, err := h.services.ParseToken(event.AuthToken)
		if err != nil {
			return err
		}
		mes.UserIdFrom = userId
		h.services.Messages.Create(mes)
		users, err := h.services.Conversation.GetAllUsersAtConv(mes.ConversationIdTo)
		for _, usr := range users {
			fmt.Println(usr.UserId)

			client, ok := h.clients[usr.UserId]
			if !ok {
				fmt.Print("notok")
				continue
			}
			client.send <- []byte(mes.Text)
			select {
			case client.send <- []byte(mes.Text):

			default:
				close(client.send)
				delete(h.clients, client.id)
			}
		}
	}

	return nil
}
