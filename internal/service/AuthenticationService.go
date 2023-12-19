package internal

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	model "learning-go/internal/model"
	"net/http"
)

func AuthEndpoint(c *gin.Context, db *gorm.DB) {
	var access model.AccessToken

	authToken := c.Request.Header.Get("Authorization")
	err := db.Model(model.AccessToken{Identifier: authToken}).Find(&access).Error

	if err != nil || authToken == "" {
		http.Error(c.Writer, "Access Denied", http.StatusForbidden)
		return
	}
}
