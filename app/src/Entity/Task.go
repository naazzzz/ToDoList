package Entity

import "gorm.io/gorm"

type TaskDTO struct {
	Title       string
	Description string
	UserId      uint
	IsFavorite  bool
}

type Task struct {
	gorm.Model
	Title       string
	Description string `gorm:"not null"`
	UserId      uint   `gorm:"not null"`
	IsFavorite  bool   `gorm:"not null"`
}
