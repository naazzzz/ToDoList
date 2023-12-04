package Service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"learning-go/src/Entity"
	"net/http"
)

func AuthEndpoint(c *gin.Context, db *gorm.DB) {
	var access Entity.AccessToken

	authToken := c.Request.Header.Get("Authorization")
	err := db.Model(Entity.AccessToken{Identifier: authToken}).Find(&access).Error

	if err != nil || authToken == "" {
		http.Error(c.Writer, "Access Denied", http.StatusForbidden)
		return
	}
}
