package models

import "gorm.io/gorm"

type Messages struct {
	gorm.Model
	FromUsersID uint   `gorm:"not null" binding:"required"`
	ToUsersID   uint   `gorm:"not null" binding:"required"`
	Subject     string `gorm:"not null" binding:"required"`
	Message     string `gorm:"not null" binding:"required"`
}
