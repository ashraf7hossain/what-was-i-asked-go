package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Body         string
	UserID       uint
	PostID       uint
	User         User          `gorm:"foreignKey:UserID"`
	Post         Post          `gorm:"foreignKey:PostID"`
	CommentVotes []CommentVote `gorm:"foreignKey:CommentID"`
}
