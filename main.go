package main

import (
	"api/web/users"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/user/{id}", users.CreateUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/user/{id}", users.GetUserHandler).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", users.DeleteUserHandler).Methods(http.MethodDelete)
	router.HandleFunc("/user/{id}", users.UpdateUserHandler).Methods(http.MethodPut)
	log.Println("Server is rUnnin' on localhost:8080")
	http.ListenAndServe(":8080", router)
}
