package service

import (
	"github.com/saparaly/forum/internal/repository"
	"github.com/saparaly/forum/models"
)

type TokenService struct {
	repo repository.Token
}

func NewTokenService(repo repository.Token) *TokenService {
	return &TokenService{repo: repo}
}

type Token interface {
	ChechUserByToken(token string) (models.User, error)
}

func (s *TokenService) ChechUserByToken(token string) (models.User, error) {
	user, err := s.repo.GetUserIDByToken(token)
	if err != nil {
		return models.User{}, err
	}
	return *user, nil
}
