package comment

import (
	"rest-in-go/initializers"
	"rest-in-go/models"
)

type CommentRepository interface {
	GetAllCommentsByPost(postID string) ([]models.Comment, error)
	PostComment(comment *models.Comment) error
}

type commentRepository struct{}

func NewCommentRepository() CommentRepository {
	return &commentRepository{}
}

func (r *commentRepository) GetAllCommentsByPost(postID string) ([]models.Comment, error) {
	var comments []models.Comment

	err := initializers.DB.Preload("User").Where("post_id = ?", postID).Find(&comments).Error

	if err != nil {
		return nil, err
	}

	return comments, nil
}


func (r *commentRepository) PostComment(comment *models.Comment) error {
	return initializers.DB.Create(comment).Error
}