package user

import (
	"fmt"
	"rest-in-go/initializers"
	"rest-in-go/models"
)

type UserRepository interface {
	CreateUser(user *models.User) (error)
	GetUserByEmail(email string) (*models.User, error)
	GetProfile(userID uint) (*models.User, error)
}

type repository struct {}

func NewUserRepository() UserRepository {
	return &repository{}
}


func (r *repository) CreateUser(user *models.User) error {
	fmt.Println("here we go from repository", user)
	return initializers.DB.Create(user).Error
}

func (r *repository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	return &user, initializers.DB.Where("email = ?", email).First(&user).Error
}

func (r *repository) GetProfile(userID uint) (*models.User, error) {
	var user models.User
	return &user, initializers.DB.First(&user, userID).Error
}
