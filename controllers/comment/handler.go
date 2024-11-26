package comment

import (
	"net/http"
	"rest-in-go/models"
	"rest-in-go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	service CommentService
}

func NewCommentHandler(service CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

func (h *CommentHandler) GetAllCommentsByPost(c *gin.Context) {
	postID := c.Param("id")

	comments, err := h.service.GetAllCommentsByPost(postID)

	if err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError) // Pass error to middleware
		return
	}

	type commentResp struct {
		CommentId uint   `json:"comment_id"`
		UserName  string `json:"user_name"`
		Body      string `json:"body"`
	}

	c.JSON(http.StatusOK, gin.H{
		"comments": utils.Map(comments, func(comment models.Comment) commentResp {
			return commentResp{
				UserName:  comment.User.Name,
				Body:      comment.Body,
				CommentId: comment.ID,
			}
		}),
	})
}

func (h *CommentHandler) PostComment(c *gin.Context) {
	var input InputComment

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(http.StatusBadRequest)
		return
	}

	userID, err := utils.GetUserID(c)
	if err != nil {
		c.Error(utils.NewError("Unauthorized")).SetMeta(http.StatusUnauthorized)
		return
	}

	comment, err := h.service.PostComment(input, userID)
	if err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, gin.H{
		"comment": comment,
	})
}

func (h *CommentHandler) UpdateComment(c *gin.Context) {
	commentID, _ := strconv.ParseUint(c.Param("comment_id"), 10, 64)
	var input struct {
		Body string `json:"body"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(http.StatusBadRequest)
		return
	}

	userID, err := utils.GetUserID(c)
	if err != nil {
		c.Error(utils.NewError("Unauthorized")).SetMeta(http.StatusUnauthorized)
		return
	}

	comment := &models.Comment{
		Body: input.Body,
	}

	err = h.service.UpdateComment(comment, uint(commentID), userID)

	if err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment updated successfully",
	})
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	commentID, _ := strconv.ParseUint(c.Param("comment_id"), 10, 64)

	userID, err := utils.GetUserID(c)
	if err != nil {
		c.Error(utils.NewError("Unauthorized")).SetMeta(http.StatusUnauthorized)
		return
	}

	err = h.service.DeleteComment(uint(commentID), userID)

	if err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment deleted successfully",
	})
}