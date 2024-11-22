package controllers

import (
	"fmt"
	"net/http"
	"rest-in-go/initializers"
	"rest-in-go/models"

	"github.com/gin-gonic/gin"
)

func PostComment(c *gin.Context) {

	var body struct {
		PostID uint   `json:"post_id"`
		Body   string `json:"body"`
	}
	
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	userID, exists := c.Get("userID")
	
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userIDUint, ok := userID.(uint)
	
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	comment := models.Comment{
		UserID: userIDUint,
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