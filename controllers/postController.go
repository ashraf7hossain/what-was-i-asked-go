package controllers

import (
	// "fmt"
	"net/http"
	"rest-in-go/initializers"
	"rest-in-go/models"
	"rest-in-go/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handle errors centrally and return a consistent response
func handleError(c *gin.Context, statusCode int, message string, details error) {
	c.JSON(statusCode, gin.H{
		"error":   message,
		"details": details.Error(),
	})
}

// Get all posts
func PostIndex(c *gin.Context) {
	var posts []models.Post

	// Fetch posts from the database
	if err := initializers.DB.Preload("Tags").Find(&posts).Error; err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to fetch posts", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

// Create a new post
func PostCreate(c *gin.Context) {
	var body struct {
		Title string   `json:"title" binding:"required"`
		Body  string   `json:"body" binding:"required"`
		Tags  []string `json:"tags" binding:"required"`
	}

	// Parse and validate the request body
	if err := c.ShouldBindJSON(&body); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid request data", err)
		return
	}

	// Extract user ID from JWT token
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Validate userID format
	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Find or create tags
	var tags []models.Tag
	for _, tagName := range body.Tags {
		var tag models.Tag
		if err := initializers.DB.FirstOrCreate(&tag, models.Tag{Name: tagName}).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Failed to process tags", err)
			return
		}
		tags = append(tags, tag)
	}

	// Create the post
	post := models.Post{
		Title:  body.Title,
		Body:   body.Body,
		UserID: userIDUint,
		Tags:   tags,
	}
	if err := initializers.DB.Create(&post).Error; err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to create post", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Post created successfully",
		"post":    post,
	})
}

// Update an existing post
func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title string   `json:"title"`
		Body  string   `json:"body"`
		Tags  []string `json:"tags"`
	}

	// Parse the request body
	if err := c.ShouldBindJSON(&body); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid request data", err)
		return
	}

	// Extract user ID from JWT token
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Validate userID format
	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Find the post
	var post models.Post
	if err := initializers.DB.Preload("Tags").First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		} else {
			handleError(c, http.StatusInternalServerError, "Failed to fetch post", err)
		}
		return
	}

	// Check if the user owns the post
	if post.UserID != userIDUint {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You are not authorized to update this post",
		})
		return
	}

	// Update the post fields
	post.Title = body.Title
	post.Body = body.Body

	// Update tags
	var newTags []models.Tag
	for _, tagName := range body.Tags {
		var tag models.Tag
		if err := initializers.DB.FirstOrCreate(&tag, models.Tag{Name: tagName}).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Failed to process tags", err)
			return
		}
		newTags = append(newTags, tag)
	}

	if err := initializers.DB.Model(&post).Association("Tags").Replace(newTags); err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to update post tags", err)
		return
	}

	// Save the updated post
	if err := initializers.DB.Save(&post).Error; err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to save updated post", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated successfully",
		"post":    post,
	})
}

func PostComments(c *gin.Context) {
	postID := c.Param("id")

	var comments []models.Comment

	err := initializers.DB.Preload("User").
		Where("post_id = ?", postID).
		Find(&comments).Error

	if err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to fetch comments", err)
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
