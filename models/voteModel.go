package models

import (
	"gorm.io/gorm"
)

type Vote struct {
	gorm.Model
	UserID uint `gorm:"not null"`
	PostID uint `gorm:"not null"`
	User   User `gorm:"foreignKey:UserID"`
	Post   Post `gorm:"foreignKey:PostID"`
	Value  int  `gorm:"not null;default:0;max:1;min:-1"`
}
