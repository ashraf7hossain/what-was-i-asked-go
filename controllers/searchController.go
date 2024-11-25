package controllers

import (
	_ "fmt"
	"net/http"
	"rest-in-go/initializers"
	"rest-in-go/models"
	"rest-in-go/utils"
	_ "rest-in-go/utils"

	"github.com/gin-gonic/gin"
)

type SearchController struct {}

func (s *SearchController) SearchByTag(c *gin.Context) {
	var body struct {
		Tags []string `json:"tags"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.Error(err).SetMeta(http.StatusBadRequest)
		return
	}

	var posts []models.Post
	// Find posts by tags
	result := initializers.DB.
		Distinct("posts.id").
		Preload("Tags"). 
		Joins("JOIN post_tags ON posts.id = post_tags.post_id").
		Joins("JOIN tags ON tags.id = post_tags.tag_id").
		Where("tags.name IN ?", body.Tags).
		Select("posts.id, posts.title, posts.body, posts.created_at").
		Find(&posts)

	if result.Error != nil {
		c.Error(result.Error).SetMeta(http.StatusInternalServerError)
		return
	}

	type resPost struct {
		Title string 		`json:"title"`
		Body  string 		`json:"body"`
		Tags  []string 	`json:"tags"`
	}

	// Map the posts to the desired response format
	c.JSON(http.StatusOK, gin.H{
		"count": len(posts),
		"posts": utils.Map(posts, func(post models.Post) resPost {
			return resPost{
				Title: post.Title,
				Body:  post.Body,
				Tags:  utils.Map(post.Tags, func(tag models.Tag) string { return tag.Name }),
			}
		}),
	})

}
