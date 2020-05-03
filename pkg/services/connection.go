package services

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var pool *sql.DB

func init() {
	var err error

	//cs := os.Getenv("CONNECTION_STRING")
	cs := "postgres://postgres:123@localhost:5432/parts?sslmode=disable"

	pool, err = sql.Open("postgres", cs)
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}
}

func getConnection() *sql.DB {
	return pool
}
