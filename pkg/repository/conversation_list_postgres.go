package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	meme "github.com/mihailco/memessenger"
)

type ConversationListPostgres struct {
	db *sqlx.DB
}

func NewConversationListPostgres(db *sqlx.DB) *ConversationListPostgres {
	return &ConversationListPostgres{db: db}
}

func (r *ConversationListPostgres) Create(userId int, info meme.ConversationStruct) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var conversationId int
	createConversationQuery := fmt.Sprintf("INSERT INTO %s (channel_username, channel_description, title, creator_id) VALUES ($1, $2, $3, $4) RETURNING ID", conversationTable)
	row := tx.QueryRow(createConversationQuery, info.ChannelName, info.Description, info.Title, info.CreatorId)
	if err := row.Scan(&conversationId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createParticipantsQuery := fmt.Sprintf("INSERT INTO %s (conversation_id, user_id) VALUES ($1, $2)", participantsTable)
	_, err = tx.Exec(createParticipantsQuery, conversationId, userId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return conversationId, tx.Commit()
}

func (r *ConversationListPostgres) GetAll(userId int) ([]meme.ConversationStruct, error) {
	var list []meme.ConversationStruct

	query := fmt.Sprintf("SELECT conv.id, conv.channel_name, conv.channel_description, conv.title FROM %s conv INNER JOIN %s part on conv.id = part.conversation_id WHERE part.user_id = $1",
		conversationTable, participantsTable)

	err := r.db.Select(&list, query, userId)

	return list, err
}
