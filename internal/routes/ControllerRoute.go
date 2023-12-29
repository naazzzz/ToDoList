package internal

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	controller "learning-go/internal/controller"
	"os"
)

func Route() *gin.Engine {
	//router := gin.New()

	switch os.Getenv("APP_ENV_MODE") {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	users := router.Group("/users")
	{
		users.POST("", controller.CreateObjUser)
		users.GET("/:id", controller.GetObjItemUser)
		users.GET("", controller.GetObjCollectionUser)
		users.PUT("/:id", controller.UpdateObjUser)
	}

	tasks := router.Group("/tasks")
	{
		tasks.POST("", controller.CreateObjTask)
		tasks.GET("/:id", controller.GetObjItemTask)
		tasks.GET("", controller.GetObjCollectionTask)
		tasks.PUT("/:id", controller.UpdateObjTask)
	}

	router.POST("/token", controller.TokenController)

	return router
}
