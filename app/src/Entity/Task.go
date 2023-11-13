package Entity

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"path"
)

type Task struct {
	gorm.Model
	Title       string
	Description string `gorm:"not null"`
	UserId      uint   `gorm:"not null"`
	IsFavorite  bool   `gorm:"not null"`
}

func (task Task) CreateObj(w http.ResponseWriter, r *http.Request, db *gorm.DB) {

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.AutoMigrate(task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Create(&task)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (task Task) GetObjItem(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	_, objId := path.Split(r.URL.Path)

	db.Model(&task).Where("id = ?", objId).Find(&task)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (task Task) GetObjCollection(w http.ResponseWriter, r *http.Request, db *gorm.DB) {

	db.Model(&task).Find(&task)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (task Task) UpdateObj(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	_, objId := path.Split(r.URL.Path)

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Model(&task).Where("id = ?", objId).Updates(&task)
	w.Header().Set("Content-Type", "application/json")
	db.Model(&task).Where("id = ?", objId).Find(&task)

	json.NewEncoder(w).Encode(task)
}
