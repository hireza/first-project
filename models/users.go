package models

import (
	"github.com/jmoiron/sqlx"
)

type Users struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *Users) CreateUser(db *sqlx.DB) error {
	query := `INSERT INTO users_collection("Name","Age") VALUES($1, $2)`
	db.QueryRow(query, u.Name, u.Age)

	return nil
}
