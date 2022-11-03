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
	byte, err := json.Marshal(u.Data)
	if err != nil {
		log.Println(err)
	}
	// other version
	// temp := []string{u.First_name, u.Last_name, u.Interests}
	// data := strings.Join(temp, " ")
	_, err = db.DB.NamedExec("INSERT INTO users (data) VALUES (:data)", map[string]interface{}{"data": byte})
	if err != nil {
		log.Println(err)
	}

	return nil
}

func (u *User) Read() (string, error) {
	var data string
	err := db.DB.Select(&data, "SELECT data FROM users where id=$1;", u.Id)
	if err != nil {
		log.Println(err)
	}

	return data, nil
}

// func (u *User) Delete() error {
// 	_, err := db.DB.NamedExec("DELETE from users where id=:id", map[string]interface{}{
// 		"id": u.Id,
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return nil
// }

// func (u *User) Update() error {
// 	_, err := db.DB.NamedExec("UPDATE users SET data =:data where id=:id", map[string]interface{}{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return nil
// }
