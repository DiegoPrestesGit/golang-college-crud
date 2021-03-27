package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type User struct {
	gorm.Model
	Id         string
	Name       string
	Email      string
	Password   string
	Created_at string
	Updated_at string
}

var UserBodyRequest struct {
	Id         string
	Name       string
	Email      string
	Password   string
	Created_at string
	Updated_at string
}

var db *gorm.DB
var err error

func Connection() {
	const (
		dialect  = "postgres"
		host     = "localhost"
		user     = "postgres"
		dbname   = "postgres"
		password = "login-password"
		port     = "5433"
	)

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, password, port)
	_, err := sql.Open(dialect, dbURI)

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	} else {
		fmt.Println("connected")
	}
	// defer db.Close()
}

func ALlUsers(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "allUsers endpoint hit")
	// defer db.Close()
}

func NewUser(response http.ResponseWriter, request *http.Request) {
	Connection()
	err := json.NewDecoder(request.Body).Decode(&UserBodyRequest)
	if nil != err {
		log.Println(err)
		return
	}

	// db.AutoMigrate(&User{})
	// db.Create(&User{})
	log.Println(UserBodyRequest)
	fmt.Fprintf(response, "%v\n", UserBodyRequest)
}

func FindUserById(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "findbyid endpoint hit")
	// defer db.Close()
}
