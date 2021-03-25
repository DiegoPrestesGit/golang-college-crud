package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Id         string
	Name       string
	Email      string
	Password   string
	Created_at string
	Updated_at string
}

func InitialMigration() {
	host := "localhost"
	user := "postgres"
	dbname := "postgres"
	password := "login-password"
	port := "5433"

	database_config := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, password, port)
	db, err := gorm.Open("postgres", database_config)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	} else {
		fmt.Println("connected")
	}
	defer db.Close()
	db.AutoMigrate(&User{})
}

func ALlUsers(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "allUsers endpoint hit")
}

func NewUser(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "newUser endpoint hit")
}

func FindUserById(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "findbyid endpoint hit")
}
