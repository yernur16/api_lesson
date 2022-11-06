package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DB *sqlx.DB

func init() {
	if err := initConfig(); err != nil {
		log.Fatalf("error with initializing configs: %s", err.Error())
	}

	conn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		viper.Get("user"),
		viper.Get("password"),
		viper.Get("host"),
		viper.Get("port"),
		viper.Get("dbname"),
		viper.Get("sslmode"))

	log.Printf("conn is %v\n", conn)
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		log.Fatalf(err.Error())
	}

	crTable := "CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY NOT NULL, data VARCHAR)"
	_, err = db.Exec(crTable)
	if err != nil {
		log.Fatalf("Error on %s", err)
	}
	DB = db
}

func initConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("development")
	return viper.ReadInConfig()
}
