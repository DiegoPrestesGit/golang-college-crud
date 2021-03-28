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
	Id       string
	Name     string
	Email    string
	Password string
}

const (
	dialect  = "postgres"
	host     = "localhost"
	user     = "postgres"
	dbname   = "postgres"
	password = "login-password"
	port     = "5433"
)

func Connection() *gorm.DB {
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

	var UserBodyRequest struct {
		Id       string
		Name     string
		Email    string
		Password string
	}

	json.NewDecoder(request.Body).Decode(&UserBodyRequest)

	createdUser := db.Create(&UserBodyRequest)

	err := createdUser.Error
	if nil != err {
		json.NewEncoder(response).Encode(err)
		return
	}
	json.NewEncoder(response).Encode(UserBodyRequest)
	defer db.Close()
}

func FindUserById(response http.ResponseWriter, request *http.Request) {
	db := Connection()
	params := mux.Vars(request)

	var user UserGolang
	rows := db.First(&user, "id = ?", params["id"])

	json.NewEncoder(response).Encode(rows)
}
