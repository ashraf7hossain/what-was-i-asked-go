package post

import (
	"rest-in-go/initializers"
	"rest-in-go/models"
	"rest-in-go/utils"
)

type PostRepository interface {
	GetAllPosts(queryParams utils.QueryParams) ([]models.Post, int, error)
	CreatePost(post *models.Post) error
	GetPostByID(postID string) (*models.Post, error)
	UpdatePost(post *models.Post) error
}

type repository struct{}

func NewPostRepository() PostRepository {
	return &repository{}

}

func (r *repository) GetAllPosts(queryParams utils.QueryParams) ([]models.Post, int, error) {
	var posts []models.Post
	var total int64

	query := initializers.DB.Model(&models.Post{}).Preload("Tags").Preload("Votes")

	// Apply search if the search parameter is provided
	if queryParams.Search != "" {
		query = query.Where("title ILIKE ?", "%"+queryParams.Search+"%")
	}

	// Count total records before applying pagination
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination and sorting
	offset := (queryParams.Page - 1) * queryParams.Limit
	if err := query.
		Order(queryParams.OrderBy).
		Limit(queryParams.Limit).
		Offset(offset).
		Find(&posts).
		Error; err != nil {
		return nil, 0, err
	}

	return posts, int(total), nil
}

func (r *repository) CreatePost(post *models.Post) error {
	return initializers.DB.Create(post).Error
}

func (r *repository) GetPostByID(postID string) (*models.Post, error) {
	var post models.Post
	if err := initializers.DB.
		Preload("Tags").
		Preload("Votes").
		First(&post, postID).
		Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *repository) UpdatePost(post *models.Post) error {
	return initializers.DB.Save(post).Error
}
