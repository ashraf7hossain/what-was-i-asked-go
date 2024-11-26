package comment

import (
	"rest-in-go/initializers"
	"rest-in-go/models"
)

type CommentRepository interface {
	GetAllCommentsByPost(postID string) ([]models.Comment, error)
	PostComment(comment *models.Comment) error
	GetCommentById(commentID uint) (*models.Comment, error)
	UpdateComment(comment *models.Comment, commentID uint) error
	DeleteComment(commentID uint) error
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

func (r *commentRepository) GetCommentById(commentID uint) (*models.Comment, error) {
	var comment models.Comment

	err := initializers.DB.Preload("User").Where("id = ?", commentID).Find(&comment).Error

	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *commentRepository) UpdateComment(comment *models.Comment, commentID uint) error {
	return initializers.DB.Model(&models.Comment{}).Where("id = ?", commentID).Updates(comment).Error
}

func (r *commentRepository) DeleteComment(commentID uint) error {
	return initializers.DB.Delete(&models.Comment{}, commentID).Error
}