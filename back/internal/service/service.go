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
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Post:          NewPostService(repos.Post),
		Token:         NewTokenService(repos.Token),
	}
}

type Comment interface{}

type Reaction interface{}

type Category interface{}
