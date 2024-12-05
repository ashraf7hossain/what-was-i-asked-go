package commentvote

import "rest-in-go/models"

type CommentVoteService interface {
	GetVotesByCommentID(commentID string) ([]*models.CommentVote, error)
	FindVoteByCommentIDAndUserID(commentID uint, userID uint) (*models.CommentVote, error)
	CreateVote(vote *models.CommentVote) error
	UpdateVote(vote *models.CommentVote) error
	DeleteVoteByCommentIDAndUserID(commentID uint, userID uint) error
}

type commentVoteService struct {
	repo CommentVoteRepository
}

func NewCommentVoteService(repo CommentVoteRepository) CommentVoteService {
	return &commentVoteService{repo: repo}
}

func (s *commentVoteService) GetVotesByCommentID(commentID string) ([]*models.CommentVote, error) {
	return s.repo.GetVotesByCommentID(commentID)
}

func (s *commentVoteService) FindVoteByCommentIDAndUserID(commentID uint, userID uint) (*models.CommentVote, error) {
	return s.repo.FindVoteByCommentIDAndUserID(commentID, userID)
}

func (s *commentVoteService) CreateVote(vote *models.CommentVote) error {
	return s.repo.CreateVote(vote)
}

func (s *commentVoteService) UpdateVote(vote *models.CommentVote) error {
	return s.repo.UpdateVote(vote)
}

func (s *commentVoteService) DeleteVoteByCommentIDAndUserID(commentID uint, userID uint) error {
	return s.repo.DeleteVoteByCommentIDAndUserID(commentID, userID)
}
