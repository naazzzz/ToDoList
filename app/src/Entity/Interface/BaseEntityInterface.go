package Interface

import (
	"gorm.io/gorm"
	"net/http"
)

type BaseEntityInterface interface {
	CreateObj(w http.ResponseWriter, r *http.Request, db *gorm.DB)
	GetObjItem(w http.ResponseWriter, r *http.Request, db *gorm.DB)
	GetObjCollection(w http.ResponseWriter, r *http.Request, db *gorm.DB)
	UpdateObj(w http.ResponseWriter, r *http.Request, db *gorm.DB)
}
