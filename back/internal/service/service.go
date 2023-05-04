package service

import (
	"github.com/saparaly/forum/internal/repository"
)

type Service struct {
	Authorization
	Post
	Comment
	Reaction
	Category
	Token
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Post:          NewPostService(repos.Post),
		Token:         NewTokenService(repos.Token),
		Comment:       NewCommentService(repos.Comment),
		User:          NewUserService(repos.User),
	}
}

type Reaction interface{}

type Category interface{}
