package service

import (
	"github.com/saparaly/forum/internal/repository"
	"github.com/saparaly/forum/models"
)

type CommentService struct {
	repo repository.Comment
}

func NewCommentService(repo repository.Comment) *CommentService {
	return &CommentService{repo: repo}
}

type Comment interface {
	CreateComment(comment models.Comment) error
	GetComment(postId int) ([]*models.Comment, error)
	GetOneComment(postId int) (models.Comment, error)
}

func (s *CommentService) CreateComment(comment models.Comment) error {
	err := s.repo.CreateComment(comment)
	if err != nil {
		return err
	}
	return nil
}

func (s *CommentService) GetComment(postId int) ([]*models.Comment, error) {
	comment, err := s.repo.GetComments(postId)
	if err != nil {
		return []*models.Comment{}, err
	}
	return comment, nil
}

func (s *CommentService) GetOneComment(postId int) (models.Comment, error) {
	comment, err := s.repo.GetOneComment(postId)
	if err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}
