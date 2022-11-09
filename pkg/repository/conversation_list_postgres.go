package repository

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	meme "github.com/mihailco/memessenger"
)

type ConversationPostgres struct {
	db *sqlx.DB
}

func NewConversationPostgres(db *sqlx.DB) *ConversationPostgres {
	return &ConversationPostgres{db: db}
}

func (r *ConversationPostgres) Create(userId int, info meme.ConversationStruct) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var conversationId int
	createConversationQuery := fmt.Sprintf("INSERT INTO %s (channel_name, channel_description, title, creator_id, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING ID", conversationTable)
	row := tx.QueryRow(createConversationQuery, info.ChannelName, info.Description, info.Title, info.CreatorId, time.Now())
	if err := row.Scan(&conversationId); err != nil {
		tx.Rollback()
		return 0, err
	}

	err = r.AddUser(userId, conversationId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return conversationId, tx.Commit()
}

func (r *ConversationPostgres) GetAll(userId int) ([]meme.ConversationStruct, error) {
	var list []meme.ConversationStruct

	query := fmt.Sprintf("SELECT conv.id, conv.channel_name, conv.channel_description, conv.title FROM %s conv INNER JOIN %s part on conv.id = part.conversation_id WHERE part.user_id=$1",
		conversationTable, participantsTable)

	err := r.db.Select(&list, query, userId)

	return list, err
}

func (r *ConversationPostgres) UpdateById(userId int, info meme.ConversationStruct) error {
	query := fmt.Sprintf("UPDATE %s SET channel_name=$1 channel_description=$2 title=$3 WHERE id=$4 AND creator_id=$5", conversationTable)
	_, err := r.db.Exec(query, info.ChannelName, info.Description, info.Title, info.Id, userId)
	return err
}

func (r *ConversationPostgres) DeleteById(userId int, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE  creator_id=$2 AND id=$1", conversationTable)
	_, err := r.db.Exec(query, userId, itemId)
	return err
}

func (r *ConversationPostgres) IsAMember(userId int, itemId int) bool {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT id FROM %s WHERE user_id=$1 AND conversation_id=$2)", participantsTable)
	r.db.Get(&exists, query, userId, itemId)
	return exists
}

func (r *ConversationPostgres) GetAllUsersAtConv(convId int) ([]meme.Userlist, error) {
	var list []meme.Userlist

	query := fmt.Sprintf("SELECT DISTINCT part.user_id FROM %s users INNER JOIN %s part on users.id = part.user_id WHERE part.conversation_id = $1 ", usersTable, participantsTable)
	err := r.db.Select(&list, query, convId)
	return list, err
}

func (r *ConversationPostgres) AddUser(userid, convId int) error {
	createParticipantsQuery := fmt.Sprintf("INSERT INTO %s (conversation_id, user_id) VALUES ($1, $2)", participantsTable)
	_, err := r.db.Exec(createParticipantsQuery, convId, userid)
	return err
}
