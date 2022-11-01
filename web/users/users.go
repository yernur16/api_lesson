package users

import (
	"api/pkg/user"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("error with mux.Vars")
	}
	us := &user.User{
		Id: id,
	}

	json.NewDecoder(r.Body).Decode(us)

	err = us.CreateUser(us.Id, us.Data)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
	}
	us := &user.User{
		Id: id,
	}
	err = us.Read(us.Id, us.Data)
	if err != nil {
		log.Println(err)
	}
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(us)

	w.WriteHeader(http.StatusOK)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
	}
	us := &user.User{
		Id: id,
	}
	err = us.Delete(us.Id, us.Data)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusNoContent)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
	}
	r.FormValue("id")

	us := &user.User{
		Id: id,
	}

	json.NewDecoder(r.Body).Decode(us)

	err = us.Update(us.Id, us.Data)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}
