package comment

import (
	"rest-in-go/models"
	"rest-in-go/utils"
)

type CommentService interface {
	GetAllCommentsByPost(postID string) ([]models.Comment, error)
	PostComment(input InputComment, userID uint) (*models.Comment, error)
	GetCommentById(commentID uint) (*models.Comment, error)
	UpdateComment(input InputComment, commentID uint, userID uint) (*models.Comment, error)
	DeleteComment(commentID uint, userID uint) error
}

type service struct {
	repo CommentRepository
}

func NewCommentService(repo CommentRepository) CommentService {
	return &service{repo: repo}
}

func (s *service) GetAllCommentsByPost(postID string) ([]models.Comment, error) {
	return s.repo.GetAllCommentsByPost(postID)
}

func (s *service) PostComment(input InputComment, userID uint) (*models.Comment, error) {
	comment := &models.Comment{
		PostID: input.PostID,
		Body  : input.Body,
		UserID: userID,
	}
	err := s.repo.PostComment(comment)

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *service) GetCommentById(commentID uint) (*models.Comment, error) {
	return s.repo.GetCommentById(commentID)
}

func (s *service) UpdateComment(input InputComment, commentID uint, userID uint) (*models.Comment, error) {
	comment, err := s.repo.GetCommentById(commentID)
	if err != nil {
		return nil, err
	}

	if comment.UserID != userID {
		return nil, utils.NewError("Unauthorized")
	}

	comment.Body = input.Body

	err = s.repo.UpdateComment(comment)

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *service) DeleteComment(commentID uint, userID uint) error {
	comment, err := s.repo.GetCommentById(commentID)
	if err != nil {
		return err
	}

	if comment.UserID != userID {
		return utils.NewError("Unauthorized")
	}
	
	return s.repo.DeleteComment(commentID)
}