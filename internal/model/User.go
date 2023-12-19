package internal

import (
	"gorm.io/gorm"
)

type UserDTO struct {
	Username string `json:"username" uri:"details" binding:"required"`
	Password string `json:"password" uri:"details" binding:"required"`
}

type User struct {
	gorm.Model
	Username string `gorm:"unique;default:null"`
	Password string `gorm:"default:null'"`
	Tasks    []Task `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
