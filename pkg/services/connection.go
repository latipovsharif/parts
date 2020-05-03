package services

import (
	"github.com/go-pg/pg/v9"
)

var pool *pg.DB

func init() {
	pool = pg.Connect(
		&pg.Options{
			Database: "parts",
			Password: "123",
			User:     "postgres",
		})
}

func getConnection() *pg.DB {
	return pool
}
