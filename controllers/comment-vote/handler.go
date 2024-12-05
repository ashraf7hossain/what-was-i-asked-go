package commentvote

import (
	"net/http"
	"rest-in-go/models"
	"rest-in-go/utils"

	"github.com/gin-gonic/gin"
)

type CommentVoteHandler struct {
	service CommentVoteService
}

func NewCommentVoteHandler(service CommentVoteService) *CommentVoteHandler {
	return &CommentVoteHandler{service: service}
}

func (h *CommentVoteHandler) GetVotesByCommentID(c *gin.Context) {
	commentID := c.Param("commentID")
	votes, err := h.service.GetVotesByCommentID(commentID)
	if err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}

	count := len(votes)

	response := map[string]interface{}{
		"message": "Successfully retrieved votes",
		"count":   count,
		"voters": utils.Map(votes, func(vote *models.CommentVote) map[string]interface{} {
			return map[string]interface{}{
				"username": vote.User.Name,
				"id":       vote.UserID,
			}
		}),
	}
	c.JSON(http.StatusOK, gin.H{"votes": response})
}

func (h *CommentVoteHandler) CreateVote(c *gin.Context) {
	var input InputCommentVote

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

	vote, err := h.service.FindVoteByCommentIDAndUserID(input.CommentID, userID)

	// never voted at all
	if err != nil {
		vote = &models.CommentVote{
			UserID:    userID,
			CommentID: input.CommentID,
			Value:     input.Value,
		}
		err = h.service.CreateVote(vote)
		if err != nil {
			c.Error(err).SetMeta(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Vote created successfully",
		})
		return
	}

	// delete vote if same vote is cast
	if vote.Value == input.Value {
		err = h.service.DeleteVoteByCommentIDAndUserID(input.CommentID, userID)
		if err != nil {
			c.Error(err).SetMeta(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Vote deleted successfully",
		})
		return
	}

	// update vote if different vote is cast
	if vote.Value != input.Value {
		vote.Value = input.Value
		err = h.service.UpdateVote(vote)
		if err != nil {
			c.Error(err).SetMeta(http.StatusInternalServerError)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Vote already exists",
		"vote":    vote,
	})

}
