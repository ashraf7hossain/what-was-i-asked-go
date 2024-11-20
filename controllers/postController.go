package controllers

import (
	_"fmt"
	"rest-in-go/initializers"
	"rest-in-go/models"

	"github.com/gin-gonic/gin"
	_"gorm.io/gorm"
)

func PostIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostCreate(c *gin.Context) {
	// Define the request body structure
	var body struct {
		Title string   `json:"title" binding:"required"`
		Body  string   `json:"body" binding:"required"`
		Tags  []string `json:"tags" binding:"required"`
	}

	// Parse and validate the JSON request body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request data",
		})
		return
	}

	// Extract the user ID from the JWT token
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	// Ensure userID is of the correct type (uint)
	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(400, gin.H{
			"error": "Invalid user ID format",
		})
		return
	}

	// find or create new tags

	var tags []models.Tag

	for _, tagName := range body.Tags {
		var tag models.Tag
		initializers.DB.FirstOrCreate(&tag, models.Tag{Name: tagName})
		tags = append(tags, tag)
	}

	// Create the post object with the parsed data and the user ID from the token
	post := models.Post{
		Title:  body.Title,
		Body:   body.Body,
		UserID: userIDUint, // Store the user ID directly in the Post model
		Tags:   tags,
	}

	// Save the post to the database
	if err := initializers.DB.Create(&post).Error; err != nil {
		c.JSON(500, gin.H{
			"error":   "Failed to create post",
			"details": err.Error(),
		})
		return
	}

	response := gin.H{
		"title": post.Title,
		"body":  post.Body,
		"tags":  post.Tags,
	}

	// Respond with success
	c.JSON(201, gin.H{
		"message": "Post created successfully",
		"post":    response,
	})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	var post models.Post

	initializers.DB.First(&post, id)

	post.Title = body.Title
	post.Body = body.Body

	initializers.DB.Save(&post)

	c.JSON(200, gin.H{
		"message": "Post updated successfully",
		"post":    post,
	})
}
