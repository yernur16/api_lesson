package user

import (
	"api/db"
	"encoding/json"
	"log"
)

type User struct {
	Id   int
	Data `json:"data"`
}

type Data struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Interests  string `json:"interests"`
}

func (u *User) CreateUser() error {
	bytes, err := json.Marshal(u.Data)
	if err != nil {
		log.Println(err)
	}

	_, err = db.DB.NamedExec("INSERT INTO users (data) VALUES (:data)", map[string]interface{}{
		"data": bytes,
	})
	if err != nil {
		log.Println(err)
	}

	return nil
}

func (u *User) Read() (string, error) {
	var data string
	var err error
	row := db.DB.QueryRow("SELECT data FROM users where id=$1", u.Id)
	row.Scan(&data)
	if err != nil {
		log.Println(err)
	}

	return data, nil
}

func (u *User) Delete() error {
	_, err := db.DB.NamedExec("DELETE from users where id=:id", map[string]interface{}{
		"id": u.Id,
	})
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (u *User) Update() error {
	bytes, err := json.Marshal(u.Data)
	if err != nil {
		log.Println(err)
	}
	_, err = db.DB.NamedExec("UPDATE users SET data =:data where id=:id", map[string]interface{}{
		"id":   u.Id,
		"data": bytes,
	})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
