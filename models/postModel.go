package models


import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string
	Body     string
	UserID   uint
	User     User 		`gorm:"foreignKey:UserID"`
	Comments []Comment
	Tags     []Tag 		`gorm:"many2many:post_tags;"`
	Votes    []Vote 	`gorm:"foreignKey:PostID"`
}