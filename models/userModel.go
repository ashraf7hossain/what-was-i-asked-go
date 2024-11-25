package models

import (
	"errors"
	"regexp"

	"gorm.io/gorm"
)
 
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique;not null;validate:?"`
	Password string
	Role     string `gorm:"default:user"`
	Posts    []Post `gorm:"foreignKey:UserID"`
}

// BeforeSave is a GORM hook that validates the user before saving to the database
func (u *User) BeforeSave(tx *gorm.DB) error {
	// Regular expression for email validation
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`)

	if !emailRegex.MatchString(u.Email) {
		return errors.New("invalid email address")
	}

	return nil
}