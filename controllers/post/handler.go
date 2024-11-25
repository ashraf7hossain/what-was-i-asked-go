package post

import (
	"net/http"
	// "rest-in-go/post"
	"rest-in-go/utils"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service PostService
}

func NewPostHandler(service PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	posts, err := h.service.GetPosts()
	if err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var input InputPost

	// Validate the request body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(http.StatusBadRequest)
		return
	}

	userID, err := utils.GetUserID(c)
	if err != nil {
		c.Error(utils.NewError("Unauthorized")).SetMeta(http.StatusUnauthorized)
		return
	}

	createdPost, err := h.service.CreateNewPost(input, userID)
	if err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully", "post": createdPost})
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	var input InputPost
	postID := c.Param("id")

	// Validate the request body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(http.StatusBadRequest)
		return
	}

	userID, err := utils.GetUserID(c)
	if err != nil {
		c.Error(utils.NewError("Unauthorized")).SetMeta(http.StatusUnauthorized)
		return
	}

	updatedPost, err := h.service.UpdateExistingPost(postID, input, userID)
	if err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully", "post": updatedPost})
}
