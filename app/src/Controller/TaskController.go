package Controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"learning-go/src/Entity"
	"learning-go/src/Service"
	"net/http"
)

// CreateObjTask godoc
// @Summary Create a new Task entity
// @tags Tasks
// @description Create a new Task entity
// @Accept json
// @Produce json
// @Param input body Entity.TaskDTO true "Task info"
// @Router /tasks [post]
func CreateObjTask(c *gin.Context) {
	var task Entity.Task

	db := Service.CreateConnection()

	err := json.NewDecoder(c.Request.Body).Decode(&task)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.AutoMigrate(task)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	db.Create(&task)

	c.Writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(c.Writer).Encode(task)
	if err != nil {
		return
	}
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
	var task Entity.Task

	db := Service.CreateConnection()

	objId := c.Param("id")

	err := db.Where("id = ?", objId).First(&task).Error

	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	c.Writer.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(c.Writer).Encode(task)
	if err != nil {
		return
	}

}

// GetObjCollectionTask godoc
// @Summary Get Object Collection Task entity
// @tags Tasks
// @description Get Object Collection Task entity
// @Accept json
// @Produce json
// @Router /tasks [get]
func GetObjCollectionTask(c *gin.Context) {
	var tasks []Entity.Task

	db := Service.CreateConnection()

	db.Model(&tasks).Find(&tasks)

	c.Writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(c.Writer).Encode(tasks)
	if err != nil {
		return
	}
}

// UpdateObjTask godoc
// @Summary Update Object Task entity
// @tags Tasks
// @description Update Object Task entity
// @Accept json
// @Produce json
// @Param id   path      int  true  "Task ID"
// @Router /tasks/{id} [put]
func UpdateObjTask(c *gin.Context) {
	var task Entity.Task

	db := Service.CreateConnection()

	objId := c.Param("id")

	err := json.NewDecoder(c.Request.Body).Decode(&task)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.Where("id = ?", objId).First(&task).Updates(&task).Error
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	db.Where("id = ?", objId).First(&task)

	err = json.NewEncoder(c.Writer).Encode(task)
	if err != nil {
		return
	}
}
