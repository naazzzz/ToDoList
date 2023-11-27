package Entity

import (
	"gorm.io/gorm"
)

type UserDTO struct {
	Username string
	Password string
}

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Tasks    []Task `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
