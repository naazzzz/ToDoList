package Service

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learning-go/src/Entity"
	"os"
)

func CreateConnection() *gorm.DB {
	path, _ := os.LookupEnv("DATABASE_URL")

	db, err := gorm.Open(mysql.Open(path), &gorm.Config{})

	if err != nil {
		fmt.Errorf("db errors: %w", err)
	}

	err = db.AutoMigrate(Entity.User{})
	err = db.AutoMigrate(Entity.Task{})

	return db
}
