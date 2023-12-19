// main.go
package main

import (
	"github.com/joho/godotenv"
	_ "learning-go/api"
	routes "learning-go/internal/routes"
	"log"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

///@title TODOList API
//@version 1.0
//@description API Server for TodoList Application

//@host localhost:8080
//@BasePath /

// @securityDefinitions.oauth2.password oauth
// @tokenUrl /token
// @in header
// @name Authorization
func main() {

	routes.Route()
}
