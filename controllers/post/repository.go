package post

import (
	"rest-in-go/initializers"
	"rest-in-go/models"
)

type PostRepository interface {
	GetAllPosts() ([]models.Post, error)
	CreatePost(post *models.Post) error
	GetPostByID(postID string) (*models.Post, error)
	UpdatePost(post *models.Post) error
}

type repository struct{}

func NewPostRepository() PostRepository {
	return &repository{}
}

func (r *repository) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	if err := initializers.DB.Preload("Tags").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *repository) CreatePost(post *models.Post) error {
	return initializers.DB.Create(post).Error
}

func (r *repository) GetPostByID(postID string) (*models.Post, error) {
	var post models.Post
	if err := initializers.DB.Preload("Tags").First(&post, postID).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *repository) UpdatePost(post *models.Post) error {
	return initializers.DB.Save(post).Error
}
