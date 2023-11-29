package Controller

import (
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
	var tasks []Entity.Task

	db := Service.CreateConnection()

	if err := c.BindJSON(&user); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Create(&user); err.Error != nil {
		http.Error(c.Writer, err.Error.Error(), http.StatusBadRequest)
		return
	}
	user.Tasks = tasks

	c.JSON(200, &user)
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

	c.JSON(200, &user)
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
	var tasks []Entity.Task

	db := Service.CreateConnection()
	db.Model(&Entity.User{}).Find(&users)

	for index, element := range users {
		db.Model(Entity.Task{}).Where("user_id = ?", element.ID).Find(&tasks)
		users[index].Tasks = tasks
	}

	c.JSON(200, &users)
}

// UpdateObjUser godoc
// @Summary Update Object User entity
// @tags Users
// @description Update Object User entity
// @Accept json
// @Produce json
// @Param id   path      int  true  "User ID"
// @Param input body Entity.UserDTO true "User info"
// @Router /users/{id} [put]
func UpdateObjUser(c *gin.Context) {
	var user Entity.User

	db := Service.CreateConnection()

	objId := c.Param("id")

	if err := c.BindJSON(&user); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	err := db.Where("id = ?", objId).Updates(&Entity.User{Username: user.Username, Password: user.Password}).Error

	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	db.Model(&user).Where("id = ?", objId).Find(&user)

	c.JSON(200, &user)
}
