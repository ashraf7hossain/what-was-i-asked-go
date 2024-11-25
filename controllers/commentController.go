package controllers

import (
	"fmt"
	"net/http"
	"rest-in-go/initializers"
	"rest-in-go/models"
	"rest-in-go/utils"

	"github.com/gin-gonic/gin"
)

type Comment struct {}

func (cmt *Comment)PostComment(c *gin.Context) {

	var body struct {
		PostID uint   `json:"post_id"`
		Body   string `json:"body"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	userID, err := utils.GetUserID(c)
	if err != nil {
		c.Error(utils.NewError("Unauthorized")).SetMeta(http.StatusUnauthorized)
		return
	}

	comment := models.Comment{
		UserID: userID,
		PostID: body.PostID,
		Body:   body.Body,
	}

	if err := initializers.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Comment added successfully to %d", body.PostID),
		"comment": comment,
	})

}
