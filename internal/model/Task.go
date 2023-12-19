package internal

import "gorm.io/gorm"

type TaskDTO struct {
	Title       string `json:"title" uri:"details" binding:"required"`
	Description string `json:"description" uri:"details" binding:"required"`
	UserId      uint   `json:"userId" uri:"details" binding:"required"`
	IsFavorite  bool   `json:"isFavorite" uri:"details" binding:"required"`
}

type Task struct {
	gorm.Model
	Title       string `gorm:"default:null"`
	Description string `gorm:"default:null"`
	UserId      uint   `gorm:"default:null"`
	IsFavorite  bool   `gorm:"default:null"`
}
