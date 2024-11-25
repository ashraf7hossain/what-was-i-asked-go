package user

import (
	"rest-in-go/models"
)

type UserService interface {
	CreateUser(input InputRegisterUser) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetProfile(userID uint) (*models.User, error)
}

type service struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &service{repo: repo}
}

func (s *service) CreateUser(input InputRegisterUser) (*models.User, error) {
	user := &models.User{
		Name: input.Name, 
		Email: input.Email, 
		Password: input.Password, 
		Role: "user",
	}

	err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) GetProfile(userID uint) (*models.User, error) {
	profile, err := s.repo.GetProfile(userID)
	if err != nil {
		return nil, err
	}
	return profile, nil
}