package internal

import (
	"github.com/gin-gonic/gin"
	model "learning-go/internal/model"
	service "learning-go/internal/service"
	"net/http"
)

// CreateObjTask godoc
// @Summary Create a new Task entity
// @tags Tasks
// @description Create a new Task entity
// @Accept json
// @Produce json
// @Param input body model.TaskDTO true "Task info"
// @Router /tasks [post]
func CreateObjTask(c *gin.Context) {
	var task model.Task

	db := service.CreateConnection()

	if err := c.BindJSON(&task); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	db.Create(&task)

	c.JSON(200, task)
}

// GetObjItemTask godoc
// @Summary Get Object item Task entity
// @tags Tasks
// @description Get Object item Task entity
// @Accept json
// @Produce json
// @Param id   path      int  true  "Task ID"
// @Router /tasks/{id} [get]
func GetObjItemTask(c *gin.Context) {
	var task model.Task

	db := service.CreateConnection()

	objId := c.Param("id")

	err := db.Where("id = ?", objId).First(&task).Error

	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	c.JSON(200, task)
}

// GetObjCollectionTask godoc
// @Summary Get Object Collection Task entity
// @tags Tasks
// @description Get Object Collection Task entity
// @Accept json
// @Produce json
// @Router /tasks [get]
func GetObjCollectionTask(c *gin.Context) {
	var tasks []model.Task

	db := service.CreateConnection()

	db.Model(&tasks).Find(&tasks)

	c.JSON(200, tasks)
}

// UpdateObjTask godoc
// @Summary Update Object Task entity
// @tags Tasks
// @description Update Object Task entity
// @Accept json
// @Produce json
// @Param id   path      int  true  "Task ID"
// @Param input body model.TaskDTO true "Task info"
// @Router /tasks/{id} [put]
func UpdateObjTask(c *gin.Context) {
	var task model.Task

	db := service.CreateConnection()

	objId := c.Param("id")

	if err := c.BindJSON(&task); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	err := db.Where("id = ?", objId).First(&task).Updates(&task).Error
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	db.Where("id = ?", objId).First(&task)

	c.JSON(200, task)
}
