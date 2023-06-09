package service

import (
	"github.com/saparaly/forum/internal/repository"
	"github.com/saparaly/forum/models"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

type User interface {
	GetUsers() ([]models.User, error)
	AddUser(username string, currectUserId int) error
	GetUserIdByUsername(username string) (int, error)
	UpdateUser(user models.User) error
	GetFollowedUsersPost(currectUser models.User) ([]models.Post, error)
	GetUserCreatedPosts(id int) ([]models.Post, error)
	GetLikedPosts(id int) ([]models.Post, error)
}

func (s *UserService) GetUsers() ([]models.User, error) {
	users, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) AddUser(username string, currectUserId int) error {
	err := s.repo.AddUser(username, currectUserId)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserIdByUsername(username string) (int, error) {
	id, err := s.repo.GetUserIdByUsername(username)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *UserService) GetUserById(id int) (models.User, error) {
	user, err := s.repo.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(user models.User) error {
	err := s.repo.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetFollowedUsersPost(currectUser models.User) ([]models.Post, error) {
	post, err := s.repo.GetFollowedUsersPost(currectUser)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *UserService) GetUserCreatedPosts(id int) ([]models.Post, error) {
	posts, err := s.repo.GetUserCreatedPosts(id)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *UserService) GetLikedPosts(id int) ([]models.Post, error) {
	posts, err := s.repo.GetLikedPosts(id)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
