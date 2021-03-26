package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/golang/users", ALlUsers).Methods("GET")
	router.HandleFunc("/api/golang/users/{id}", FindUserById).Methods("GET")
	router.HandleFunc("/api/golang/users", NewUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":5666", router))
}

func main() {
	fmt.Println("Go!")
	handleRequests()
}
