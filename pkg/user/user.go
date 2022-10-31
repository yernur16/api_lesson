package user

import (
	"log"

	// _ "github.com/mattn/go-sqlite3"
	"api/db"
)

type User struct {
	Id   int
	Data string `json:"data"`
}

func NewUser(id int, data string) *User {
	return &User{
		Id:   id,
		Data: data,
	}
}

// func DB() (*sql.DB, error) {
// 	db, err := sql.Open("sqlite3", "api.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	crTable := "CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY, username TEXT, password TEXT)"
// 	_, err = db.Exec(crTable)
// 	if err != nil {
// 		log.Fatalf("Error on %s", err)
// 	}
// 	return db, nil
// }

func (u *User) CreateUser(id int, data string) error {
	_, err := db.DB.NamedExec("INSERT INTO users (id, data) VALUES (:id, :data)", map[string]interface{}{
		"id":   u.Id,
		"data": u.Data,
	})
	if err != nil {
		log.Println(err)
	}

	return nil
}

func (u *User) Read(id int, data string) error {
	rows, err := db.DB.NamedQuery("SELECT id, data FROM users where id=:id", map[string]interface{}{
		"id": u.Id,
	})
	for rows.Next() {
		err := rows.StructScan(u)
		if err != nil {
			log.Println(err)
		}
	}

	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (u *User) Delete(id int, data string) error {
	_, err := db.DB.NamedExec("DELETE from users where id=:id", map[string]interface{}{
		"id": u.Id,
	})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (u *User) Update(id int, data string) error {
	_, err := db.DB.NamedExec("UPDATE users SET data =:data where id=:id", map[string]interface{}{
		"id":   u.Id,
		"data": u.Data,
	})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
