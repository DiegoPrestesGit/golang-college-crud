package main

import (
	"fmt"
	"net/http"
)

func ALlUsers(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "allUsers endpoint hit")
}

func NewUser(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "newUser endpoint hit")
}

func FindUserById(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "findbyid endpoint hit")
}
