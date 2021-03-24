package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	id         string
	name       string
	email      string
	password   string
	created_at string
	updated_at string
}

func InitialMigration() {
	database_config := "host:localhost user=postgres password=login-password dbname=postgres port=5433 sslmode=disable"
	db, err := gorm.Open(postgres., &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()
	db.AutoMigrate((&User{}))
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
