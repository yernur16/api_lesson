package user

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id       int
	Username string `json:"username"`
	Password string `json:"password"`
}

func DB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	crTable := "CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY, username TEXT, password TEXT)"
	_, err = db.Exec(crTable)
	if err != nil {
		log.Fatalf("Error on %s", err)
	}
	return db, nil
}

func (u *User) CreateUser(db *sql.DB) error {
	query := `INSERT INTO users (username, password) VALUES ($1, $2)`
	_, err := db.Exec(query, u.Username, u.Password)
	if err != nil {
		log.Println(err)
	}

	return nil

}

func (u *User) Read(db *sql.DB) error {
	query := `SELECT username, password FROM users where id=$1`
	err := db.QueryRow(query, u.Id).Scan(&u.Username, &u.Password)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (u *User) Delete(db *sql.DB) error {
	query := `DELETE from users where id=$1`
	_, err := db.Exec(query, u.Id)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
