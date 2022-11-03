package users

import (
	"api/pkg/user"
	"encoding/json"
	"io/ioutil"
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

	bb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(bb, us)
	if err != nil {
		log.Println(err)
	}

	// Alternative version
	// json.NewDecoder(r.Body).Decode(us)

	err = us.CreateUser()
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
	data, err := us.Read()
	if err != nil {
		log.Println(err)
	}

	w.Header().Add("Content-type", "application/json")

	w.Write([]byte(data))

	w.WriteHeader(http.StatusOK)
}

// func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(mux.Vars(r)["id"])
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	us := &user.User{
// 		Id: id,
// 	}
// 	err = us.Delete(us.Id, us.Data)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	w.WriteHeader(http.StatusNoContent)
// }

// func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(mux.Vars(r)["id"])
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	r.FormValue("id")

// 	us := &user.User{
// 		Id: id,
// 	}

// 	json.NewDecoder(r.Body).Decode(us)

// 	err = us.Update(us.Id, us.Data)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	w.WriteHeader(http.StatusOK)
// }
