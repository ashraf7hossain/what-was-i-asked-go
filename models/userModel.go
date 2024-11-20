package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     	string
	Email    	string
	Password 	string
	Role     	string `gorm:"default:user"`
	Posts    	[]Post `gorm:"foreignKey:UserID"`
}