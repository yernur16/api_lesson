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

	conn := fmt.Sprintf("user=%s dbname=%s host=%s port=%s password=%s sslmode=%s",
		viper.Get("services.postgres.configs.0"),
		viper.Get("services.postgres.configs.1"),
		viper.Get("services.postgres.configs.2"),
		viper.Get("services.postgres.configs.3"),
		viper.Get("services.postgres.configs.4"),
		viper.Get("services.postgres.configs.5"))

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
	viper.AddConfigPath(".")
	viper.SetConfigName("docker-compose")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
