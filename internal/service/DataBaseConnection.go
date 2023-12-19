package internal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	model "learning-go/internal/model"
	"os"
)

func CreateConnection() *gorm.DB {
	path, _ := os.LookupEnv("DATABASE_URL")

	db, err := gorm.Open(mysql.Open(path), &gorm.Config{})

	if err != nil {
		fmt.Errorf("db errors: %w", err)
	}

	err = db.AutoMigrate(model.User{}, model.Task{}, model.AccessToken{}, model.Client{}, model.RefreshToken{})

	return db
}
