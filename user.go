package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type UserGolang struct {
	gorm.Model
	Id       int `gorm:"typevarchar(100);unique_index"`
	Name     string
	Email    string `gorm:"typevarchar(100);unique_index"`
	Password string
}

func Connection() *gorm.DB {
	const (
		dialect  = "postgres"
		host     = "localhost"
		user     = "postgres"
		dbname   = "postgres"
		password = "login-password"
		port     = "5433"
	)

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, password, port)
	db, err := gorm.Open(dialect, dbURI)

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	} else {
		fmt.Println("connected")
	}
	db.AutoMigrate(&UserGolang{})
	return db
}

func ALlUsers(response http.ResponseWriter, request *http.Request) {
	db := Connection()

	var users []UserGolang
	db.Find(&users)
	json.NewEncoder(response).Encode(&users)

	defer db.Close()
}

func NewUser(response http.ResponseWriter, request *http.Request) {
	db := Connection()

	var createUser UserGolang

	json.NewDecoder(request.Body).Decode(&createUser)

	createdUser := db.Create(&createUser)

	err := createdUser.Error
	if nil != err {
		json.NewEncoder(response).Encode(&createUser)
		return
	}

	json.NewEncoder(response).Encode(createUser)
	defer db.Close()
}

func FindUserById(response http.ResponseWriter, request *http.Request) {
	db := Connection()
	params := mux.Vars(request)

	var user UserGolang
	rows := db.First(&user, "id = ?", params["id"])

	json.NewEncoder(response).Encode(rows)
	defer db.Close()
}
