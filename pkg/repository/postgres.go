package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	usersTable        = "users"
	messagesTable     = "messages"
	conversationTable = "conversation"
	participantsTable = "participants"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLmode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLmode))

	println(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLmode))

	if err != nil {
		fmt.Printf("sqlx open error")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Printf("ping error")
		return nil, err
	}
	return db, nil
}
