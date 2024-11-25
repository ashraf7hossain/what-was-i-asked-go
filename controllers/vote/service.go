package vote

import (
	"rest-in-go/models"
)

type VoteService interface {
	GetVotesByPostID(postID string) ([]*models.Vote, error)
	CreateVote(vote *models.Vote) error
	UpdateVote(vote *models.Vote) error
	FindVoteByPostIDAndUserID(postID uint, userID uint) (*models.Vote, error)
	DeleteVoteByPostIDAndUserID(postID uint, userID uint) (error)
}

type voteService struct {
	repo VoteRepository
}

func NewVoteService(repo VoteRepository) VoteService {
	return &voteService{repo: repo}
}

func (s *voteService) GetVotesByPostID(postID string) ([]*models.Vote, error) {
	return s.repo.GetVotesByPostID(postID)
}

func (s *voteService) CreateVote(vote *models.Vote) error {
	return s.repo.CreateVote(vote)
}

func (s *voteService) FindVoteByPostIDAndUserID(postID uint, userID uint) (*models.Vote, error) {
	return s.repo.FindVoteByPostIDAndUserID(postID, userID)
}

func (s *voteService) UpdateVote(vote *models.Vote) error {
	return s.repo.UpdateVote(vote)
}

func (s *voteService) DeleteVoteByPostIDAndUserID(postID uint, userID uint) error{
	return s.repo.DeleteVoteByPostIDAndUserID(postID, userID)
}