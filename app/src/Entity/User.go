package Entity

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"path"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Tasks    []Task `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (user User) CreateObj(w http.ResponseWriter, r *http.Request, db *gorm.DB) {

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Вынести
	err = db.AutoMigrate(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Create(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (user User) GetObjItem(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var tasks []Task

	_, objId := path.Split(r.URL.Path)
	db.Model(&user).Where("id = ?", objId).Find(&user)
	db.Model(Task{UserId: user.ID}).Find(&tasks)
	user.Tasks = tasks

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (user User) GetObjCollection(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var users []User
	db.Model(&user).Find(&users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (user User) UpdateObj(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	_, objId := path.Split(r.URL.Path)

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Model(&user).Where("id = ?", objId).Updates(&user)
	w.Header().Set("Content-Type", "application/json")
	db.Model(&user).Where("id = ?", objId).Find(&user)

	json.NewEncoder(w).Encode(user)
}
