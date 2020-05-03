package services

import (
	"os"

	"github.com/go-pg/pg/v9"
)

var pool *pg.DB

func init() {
	pool = pg.Connect(
		&pg.Options{
			Addr:       "postgres:5432",
			Database:   "parts",
			Password:   os.Getenv("DATABASE_PASS"),
			User:       "postgres",
			MaxRetries: 100,
		})
}

func getConnection() *pg.DB {
	return pool
}
