package repository

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	meme "github.com/mihailco/memessenger"
)

type MessagesPostgres struct {
	db *sqlx.DB
}

func NewMessagesPostgres(db *sqlx.DB) *MessagesPostgres {
	return &MessagesPostgres{db: db}
}

func (r *MessagesPostgres) Create(info meme.MessageList) (int, error) {
	id := 0
	querry := fmt.Sprintf("INSERT INTO %s (message_text, user_id_from, conversation_id_to, typeofmessage, createdat) values ($1, $2, $3, $4, $5) RETURNING id", messagesTable)
	row := r.db.QueryRow(querry, info.Text, info.UserIdFrom, info.ConversationIdTo, info.TypeOfMessage, time.Now())
	if err := row.Scan(&id); err != nil {
		fmt.Println(err)

		return 0, err
	}

	return id, nil
}

func (r *MessagesPostgres) Read(userId int, conversationId int, loadFromMessageId int) ([]meme.MessageList, error) {

	exists := false
	querry1 := fmt.Sprintf("SELECT EXISTS(SELECT id FROM %s WHERE conversation_id=$1 AND user_id=$2)", participantsTable)
	if err := r.db.Get(&exists, querry1, conversationId, userId); err != nil {
		return nil, err
	}

	var info []meme.MessageList
	querry2 := fmt.Sprintf("SELECT * FROM %s WHERE id > $1 and conversation_id_to = $2 LIMIT 100", messagesTable)
	if err := r.db.Select(&info, querry2, userId, conversationId); err != nil {
		return nil, err
	}
	return info, nil
}
