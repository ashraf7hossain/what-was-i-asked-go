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

// Get all posts
func PostIndex(c *gin.Context) {
	var posts []models.Post

	// Fetch posts from the database
	if err := initializers.DB.Preload("Tags").Find(&posts).Error; err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError) // Pass error to middleware
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

// Create a new post
func CreatePost(c *gin.Context) {
	var body struct {
		Title string   `json:"title" binding:"required"`
		Body  string   `json:"body" binding:"required"`
		Tags  []string `json:"tags" binding:"required"`
	}

	// Parse and validate the request body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err).SetMeta(http.StatusBadRequest)
		return
	}

	userID, err := utils.GetUserID(c)
	if err != nil {
		c.Error(utils.NewError("Unauthorized")).SetMeta(http.StatusUnauthorized)
		return
	}
	// Find or create tags
	var tags []models.Tag
	for _, tagName := range body.Tags {
		var tag models.Tag
		if err := initializers.DB.FirstOrCreate(&tag, models.Tag{Name: tagName}).Error; err != nil {
			c.Error(err).SetMeta(http.StatusInternalServerError)
			return
		}
		tags = append(tags, tag)
	}

	// Create the post
	post := models.Post{
		Title:  body.Title,
		Body:   body.Body,
		UserID: userID,
		Tags:   tags,
	}
	if err := initializers.DB.Create(&post).Error; err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Post created successfully",
		"post":    post,
	})
}

// Update an existing post
func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title string   `json:"title"`
		Body  string   `json:"body"`
		Tags  []string `json:"tags"`
	}

	// Parse the request body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err).SetMeta(http.StatusBadRequest)
		return
	}

	userID, err := utils.GetUserID(c)
	if err != nil {
		c.Error(utils.NewError("Unauthorized")).SetMeta(http.StatusUnauthorized)
		return
	}
	// Find the post
	var post models.Post
	if err := initializers.DB.Preload("Tags").First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Error(err).SetMeta(http.StatusNotFound)
		} else {
			c.Error(err).SetMeta(http.StatusInternalServerError)
		}
		return
	}

	// Check if the user owns the post
	if post.UserID != userID {
		c.Error(utils.NewError("Unauthorized")).SetMeta(http.StatusForbidden)
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
			c.Error(err).SetMeta(http.StatusInternalServerError)
			return
		}
		newTags = append(newTags, tag)
	}

	if err := initializers.DB.Model(&post).Association("Tags").Replace(newTags); err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}

	// Save the updated post
	if err := initializers.DB.Save(&post).Error; err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated successfully",
		"post":    post,
	})
}

// Get all comments of a post
func GetComments(c *gin.Context) {
	postID := c.Param("id")

	var comments []models.Comment

	err := initializers.DB.Preload("User").
		Where("post_id = ?", postID).
		Find(&comments).Error

	if err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
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
