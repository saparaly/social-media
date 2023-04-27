package service

import (
	"errors"

	"github.com/saparaly/forum/internal/repository"
	"github.com/saparaly/forum/models"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

type Post interface {
	CreatePost(post models.Post) (int64, error)
	GetPost(postId int) (models.Post, error)
	UpdatePost(post models.Post) error
	DeletePost(post models.Post) error
	LikePost(postId, userId int) error
	DislikePost(postId, userId int) error
}

func (s *PostService) CreatePost(post models.Post) (int64, error) {
	if post.Title == "" || post.Description == "" {
		return 0, errors.New("Title and Description fields should be fill")
	}
	// validate img in future
	post1, err := s.repo.CreatePost(post)
	if err != nil {
		return 0, err
	}
	return post1, nil
}

func (s *PostService) GetPost(postId int) (models.Post, error) {
	post, err := s.repo.GetPost(postId)
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func (s *PostService) UpdatePost(post models.Post) error {
	// some validation in future now we dont want to do that
	err := s.repo.UpdatePost(post)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) DeletePost(post models.Post) error {
	// some validation in future now we dont want to do that
	err := s.repo.DeletePost(post)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) LikePost(postId, userId int) error {
	err := s.repo.LikePost(postId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) DislikePost(postId, userId int) error {
	return nil
}
