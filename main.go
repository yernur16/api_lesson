package main

import (
	"api/web/users"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/user/{id}", users.CreateUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/user/{id}", users.GetUserHandler).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", users.DeleteUserHandler).Methods(http.MethodDelete)

	http.ListenAndServe(":8080", router)

}
