package Controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"learning-go/src/Entity"
	"learning-go/src/Service"
	"net/http"
)

// CreateObjUser godoc
// @Summary Create a new User entity
// @tags Users
// @description Create a new User entity
// @Accept json
// @Produce json
// @Param input body Entity.UserDTO true "User info"
// @Router /users [post]
func CreateObjUser(c *gin.Context) {
	var user Entity.User

	db := Service.CreateConnection()

	err := json.NewDecoder(c.Request.Body).Decode(&user)

	if err != nil {
		http.Error(c.Writer, c.Errors.String(), http.StatusBadRequest)
		return
	}
	//Вынести
	err = db.AutoMigrate(user)

	if err != nil {
		http.Error(c.Writer, c.Errors.String(), http.StatusBadRequest)
		return
	}

	db.Create(&user)

	c.Writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(c.Writer).Encode(user)
	if err != nil {
		return
	}
}

// GetObjItemUser godoc
// @Summary Get Object item User entity
// @tags Users
// @description Get Object item User entity
// @Accept json
// @Produce json
// @Param id   path      int  true  "User ID"
// @Router /users/{id} [get]
func GetObjItemUser(c *gin.Context) {
	var user Entity.User
	var tasks []Entity.Task

	db := Service.CreateConnection()

	objId := c.Param("id")

	err := db.Where("id = ?", objId).First(&user).Error

	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	db.Model(Entity.Task{UserId: user.ID}).Find(&tasks)
	user.Tasks = tasks

	c.Writer.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(c.Writer).Encode(user)
	if err != nil {
		return
	}
}

// GetObjCollectionUser godoc
// @Summary Get Object Collection User entity
// @tags Users
// @description Get Object Collection User entity
// @Accept json
// @Produce json
// @Router /users [get]
func GetObjCollectionUser(c *gin.Context) {
	var users []Entity.User

	db := Service.CreateConnection()
	db.Model(&Entity.User{}).Find(&users)

	c.Writer.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(c.Writer).Encode(users)
	if err != nil {
		return
	}
}

// UpdateObjUser godoc
// @Summary Update Object User entity
// @tags Users
// @description Update Object User entity
// @Accept json
// @Produce json
// @Param id   path      int  true  "User ID"
// @Router /users/{id} [put]
func UpdateObjUser(c *gin.Context) {
	var user Entity.User

	db := Service.CreateConnection()

	objId := c.Param("id")

	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.Where("id = ?", objId).Updates(&user).Error

	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	db.Model(&user).Where("id = ?", objId).Find(&user)

	err = json.NewEncoder(c.Writer).Encode(user)
	if err != nil {
		return
	}
}
