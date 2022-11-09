package repository

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	meme "github.com/mihailco/memessenger"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user meme.Userlist) (int, error) {
	var id int
	querry := fmt.Sprintf("INSERT INTO %s (firstname, lastname, username, password_hash, email, createdat) values ($1, $2, $3, $4, $5, $6) RETURNING id", "users")
	row := r.db.QueryRow(querry, user.Firstname, user.Lastname, user.Username, user.Password, user.Email, time.Now())
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (meme.Userlist, error) {
	var user meme.Userlist
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", "users")
	err := r.db.Get(&user, query, username, password)
	return user, err
}

func (r *AuthPostgres) Exist(id int, password string) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT id FROM %s WHERE id=$1 AND password_hash=$2)", "users")
	err := r.db.Get(&exists, query, id, password)
	return exists, err
}
