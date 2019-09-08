package models

import (
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func ConnectPosgres() *sqlx.DB {
	c, err := sqlx.Open("postgres", "postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("Unable to connect : ", err)
	}

	return c
}
