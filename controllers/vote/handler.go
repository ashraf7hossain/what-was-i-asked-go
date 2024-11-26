package vote

import (
	"net/http"
	"rest-in-go/models"
	"rest-in-go/utils"

	"github.com/gin-gonic/gin"
)

type VoteHandler struct {
	service VoteService
}

func NewVoteHandler(service VoteService) *VoteHandler {
	return &VoteHandler{service: service}
}

func (s *VoteHandler) GetVotesByPostID(c *gin.Context) {
	postID := c.Param("id")
	votes, err := s.service.GetVotesByPostID(postID)
	if err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}

	count := len(votes)

	response := map[string]interface{}{
		"message": "Successfully retrieved votes",
		"count":   count,
		"voters": utils.Map(votes, func(vote *models.Vote) map[string]interface{} {
			return map[string]interface{}{
				"username": vote.User.Name,
				"id":       vote.UserID,
			}
		}),
	}

	c.JSON(http.StatusOK, gin.H{"votes": response})
}

func (s *VoteHandler) CreateVote(c *gin.Context) {
	var input InputVote

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(http.StatusBadRequest)
		return
	}

	userID, err := utils.GetUserID(c)

	if err != nil {
		c.Error(utils.NewError("Unauthorized")).SetMeta(http.StatusUnauthorized)
		return
	}

	if input.Value != 1 && input.Value != -1 {
		c.Error(utils.NewError("Invalid vote value")).SetMeta(http.StatusBadRequest)
		return
	}

	vote, err := s.service.FindVoteByPostIDAndUserID(input.PostID, userID)

	// never voted at all
	if err != nil {
		vote = &models.Vote{
			UserID: userID,
			PostID: input.PostID,
			Value:  input.Value,
		}
		err = s.service.CreateVote(vote)
		if err != nil {
			c.Error(err).SetMeta(http.StatusInternalServerError)
		}

		c.JSON(http.StatusOK, gin.H{"message": "Vote created successfully", "vote": vote})
		return
	}

	// delete vote if same button pressed or same vote is cast
	if input.Value == vote.Value {
		err := s.service.DeleteVoteByPostIDAndUserID(input.PostID, userID)
		if err != nil {
			c.Error(err).SetMeta(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Vote deleted successfully"})
		return
	}

	// update vote if different vote cast
	if vote.Value != input.Value {
		vote.Value = input.Value
		err = s.service.UpdateVote(vote)
		if err != nil {
			c.Error(err).SetMeta(http.StatusInternalServerError)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vote updated successfully", "vote": vote})

}
