// main.go
package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learning-go/src/Entity"
	"log"
	"net/http"
)

func main() {

	dsn := "root:root@tcp(127.0.0.1:4306)/go_base?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	var user Entity.User
	var task Entity.Task

	if err != nil {
		fmt.Errorf("db errors: %w", err)
	}

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		checkContentType(r)

		switch r.Method {
		case http.MethodPost:

			user.CreateObj(w, r, db)
			break
		case http.MethodGet:
			user.GetObjCollection(w, r, db)
			break
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		checkContentType(r)

		switch r.Method {
		case http.MethodPut:
			user.UpdateObj(w, r, db)
			break
		case http.MethodGet:
			user.GetObjItem(w, r, db)
			break
		}
	})

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		checkContentType(r)

		switch r.Method {
		case http.MethodPost:
			task.CreateObj(w, r, db)
			break
		case http.MethodGet:
			task.GetObjCollection(w, r, db)
			break
		}
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		checkContentType(r)

		switch r.Method {
		case http.MethodPost:
			task.CreateObj(w, r, db)
			break
		case http.MethodPut:
			task.UpdateObj(w, r, db)
			break
		case http.MethodGet:
			task.GetObjItem(w, r, db)
			break
		}
	})

	http.HandleFunc("/users/:id/tasks", func(w http.ResponseWriter, r *http.Request) {
		checkContentType(r)

		switch r.Method {
		case http.MethodGet:
			//Entity.GetUserItem(w, r, db)
			break
		}
	})

	err = http.ListenAndServe(":9990", nil)

	log.Fatal(err)
}

func checkContentType(r *http.Request) {
	contentType := r.Header.Get("content-type")

	if contentType != "application/json" {
		fmt.Errorf("unsupported content-type: " + contentType)
	}
}
