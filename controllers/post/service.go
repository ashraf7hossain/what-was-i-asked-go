package post

import (
	"rest-in-go/models"
	"rest-in-go/utils"
)

type PostService interface {
	GetPosts(queryParams utils.QueryParams) ([]models.Post,int, error)
	GetPostByID(postID string) (*models.Post, error)
	CreateNewPost(input InputPost, userID uint) (*models.Post, error)
	UpdateExistingPost(postID string, input InputPost, userID uint) (*models.Post, error)
	DeletePost(postID string, userID uint) (error)
}

type service struct {
	repo PostRepository
}

func NewPostService(repo PostRepository) PostService {
	return &service{repo: repo}
}

func (s *service) GetPosts(queryParams utils.QueryParams) ([]models.Post,int, error) {
	return s.repo.GetAllPosts(queryParams)
}

func (s *service) CreateNewPost(input InputPost, userID uint) (*models.Post, error) {
	// Find or create tags
	var postTags []models.Tag
	for _, tagName := range input.Tags {
		tag := models.Tag{Name: tagName}
		postTags = append(postTags, tag)
	}

	// Create a Post entity
	post := &models.Post{
		Title:  input.Title,
		Body:   input.Body,
		UserID: userID,
		Tags:   postTags,
	}

	// Save using repository
	if err := s.repo.CreatePost(post); err != nil {
		return nil, err
	}

	return post, nil
}

func (s *service) GetPostByID(postID string) (*models.Post, error) {
	return s.repo.GetPostByID(postID)
}

func (s *service) UpdateExistingPost(postID string, input InputPost, userID uint) (*models.Post, error) {
	// Fetch the existing post
	post, err := s.repo.GetPostByID(postID)
	if err != nil {
		return nil, err
	}

	// Check if the user owns the post
	if post.UserID != userID {
		return nil, utils.NewError("Unauthorized")
	}

	// Update post details
	post.Title = input.Title
	post.Body = input.Body

	// Update tags
	var updatedTags []models.Tag
	for _, tagName := range input.Tags {
		tag := models.Tag{Name: tagName}
		updatedTags = append(updatedTags, tag)
	}
	post.Tags = updatedTags

	// Save updated post
	if err := s.repo.UpdatePost(post); err != nil {
		return nil, err
	}

	return post, nil
}

func (s *service) DeletePost(postID string, userID uint) error {
	post, err := s.GetPostByID(postID)

	if err != nil {
		return err
	}

	if post.UserID != userID {
		return utils.NewError("unauthroized")
	}

	return s.repo.DeletePost(postID)
}