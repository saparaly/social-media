package repository

import (
	"database/sql"
)

type Repository struct {
	Authorization
	Post
	Comment
	Reaction
	Category
	Token
	User
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Post:          NewPostRepository(db),
		Token:         NewTokenRepository(db),
		Comment:       NewCommentRepository(db),
		User:          NewUserRepository(db),
	}
}

type Reaction interface{}

type Category interface{}
