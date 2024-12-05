package commentvote

import (
	"rest-in-go/initializers"
	"rest-in-go/models"
)

type CommentVoteRepository interface {
	CreateVote(vote *models.CommentVote) error
	UpdateVote(vote *models.CommentVote) error
	GetVotesByCommentID(commentID string) ([]*models.CommentVote, error)
	FindVoteByCommentIDAndUserID(commentID uint, userID uint) (*models.CommentVote, error)
	DeleteVoteByCommentIDAndUserID(commentID uint, userID uint) error
}

type commentVoteRepository struct{}

func NewCommentVoteRepository() CommentVoteRepository {
	return &commentVoteRepository{}
}

func (r *commentVoteRepository) CreateVote(vote *models.CommentVote) error {
	return initializers.DB.Create(vote).Error
}

func (r *commentVoteRepository) UpdateVote(vote *models.CommentVote) error {
	return initializers.DB.Save(vote).Error
}

func (r *commentVoteRepository) GetVotesByCommentID(commentID string) ([]*models.CommentVote, error) {
	var votes []*models.CommentVote

	if err := initializers.DB.
		Where("comment_id = ?", commentID).
		Find(&votes).Error; err != nil {
		return nil, err
	}

	return votes, nil
}

func (r *commentVoteRepository) FindVoteByCommentIDAndUserID(commentID uint, userID uint) (*models.CommentVote, error) {
	var vote models.CommentVote

	return &vote, initializers.DB.
	Where("comment_id = ? AND user_id = ?", commentID, userID).
	First(&vote).Error
}

func (r *commentVoteRepository) DeleteVoteByCommentIDAndUserID(commentID uint, userID uint) error {
	return initializers.DB.
	Where("comment_id = ? AND user_id = ?", commentID, userID).
	Delete(&models.CommentVote{}).
	Error
}