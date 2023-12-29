package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func Init() {
	currentWorkDirectory := os.Getenv("PWD")
	s := strings.Split(currentWorkDirectory, "/test")

	// Если отсутствует .env, берем системные переменные окружения
	if err := godotenv.Load(s[0] + "/.env.test"); err != nil {
		log.Print("No .env file found")
	}
}
