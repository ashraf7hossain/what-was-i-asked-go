package comment

import "rest-in-go/models"

type CommentService interface {
	GetAllCommentsByPost(postID string) ([]models.Comment, error)
	PostComment(input InputComment, userID uint) (*models.Comment, error)
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
