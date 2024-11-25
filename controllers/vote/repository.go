package vote

import (
	"rest-in-go/initializers"
	"rest-in-go/models"
)

type VoteRepository interface {
	CreateVote(vote *models.Vote) error
	UpdateVote(vote *models.Vote) error
	GetVotesByPostID(postID string) ([]*models.Vote, error)
	FindVoteByPostIDAndUserID(postID uint, userID uint) (*models.Vote, error)
	DeleteVoteByPostIDAndUserID(postID uint, userID uint) (error)
}

type voteRepository struct{}

func NewVoteRepository() VoteRepository {
	return &voteRepository{}
}

func (r *voteRepository) CreateVote(vote *models.Vote) error {
	return initializers.DB.Create(vote).Error
}

func (r *voteRepository) UpdateVote(vote *models.Vote) error {
	return initializers.DB.Save(vote).Error
}

func (r *voteRepository) GetVotesByPostID(postID string) ([]*models.Vote, error) {
	var votes []*models.Vote
	return votes, initializers.DB.Preload("User").Where("post_id = ?", postID).Find(&votes).Error
}

func (r *voteRepository) FindVoteByPostIDAndUserID(postID uint, userID uint) (*models.Vote, error) {
	var vote models.Vote
	return &vote, initializers.DB.Where("post_id = ? AND user_id = ?", postID, userID).First(&vote).Error
}

func (r *voteRepository) DeleteVoteByPostIDAndUserID(postID uint, userID uint) error {
	return initializers.DB.Where("post_id = ? AND user_id = ?", postID, userID).Delete(&models.Vote{}).Error
}