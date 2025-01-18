package database

import (
	"database/sql"
	database "family/database/handlers"
	"fmt"
	"log"
	"os"
)

func NewDatabase() *database.Queries {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, name))
	if err != nil {
		log.Fatal("could not connect to database", err)
	}

	queries := database.New(db)

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	return queries
}
