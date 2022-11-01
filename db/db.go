package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	port     = "5432"
	user     = "postgres"
	password = "qwerty"
	host     = "postgres"
	dbname   = "postgres"
	sslmode  = "disable"
)

var DB *sqlx.DB

func init() {
	conn := fmt.Sprintf("user=%s dbname=%s host=%s port=%s password=%s sslmode=%s", user, dbname, host, port, password, sslmode)
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		log.Fatalf(err.Error())
	}

	crTable := "CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY, data VARCHAR)"
	_, err = db.Exec(crTable)
	if err != nil {
		log.Fatalf("Error on %s", err)
	}
	DB = db
}
