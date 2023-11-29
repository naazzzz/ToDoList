// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "learning-go/docs"
	"learning-go/src/Controller"
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

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	//вынести в init routes
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	users := router.Group("/users")
	{
		users.POST("", Controller.CreateObjUser)
		users.GET("/:id", Controller.GetObjItemUser)
		users.GET("", Controller.GetObjCollectionUser)
		users.PUT("/:id", Controller.UpdateObjUser)
	}

	tasks := router.Group("/tasks")
	{
		tasks.POST("", Controller.CreateObjTask)
		tasks.GET("/:id", Controller.GetObjItemTask)
		tasks.GET("", Controller.GetObjCollectionTask)
		tasks.PUT("/:id", Controller.UpdateObjTask)
	}

	err := router.Run()
	if err != nil {
		return
	}
}
