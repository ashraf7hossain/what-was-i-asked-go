package models

import (
	"gorm.io/gorm"
)

type CommentVote struct {
	gorm.Model
	UserID    uint    `gorm:"not null"`
	CommentID uint    `gorm:"not null"`
	User      User    `gorm:"foreignKey:UserID"`
	comment   Comment `gorm:"foreignKey:CommentID"`
	Value     int     `gorm:"not null;default:0;max:1;min:-1"`
}
