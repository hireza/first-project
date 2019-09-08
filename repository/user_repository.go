package repository

import (
	"log"

	"github.com/hireza/first-project/models"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type mysqlUser struct {
	DB *sqlx.DB
}

type Users models.Users

func GetUsers(db *sqlx.DB) ([]models.Users, error) {
	query := "SELECT * FROM users_collection"
	users := []models.Users{}

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error fetching the table")
	}

	defer rows.Close()

	for rows.Next() {
		var u models.Users
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func CreateUser(db *sqlx.DB, u models.Users) error {
	query := `INSERT INTO users_collection("Name","Age") VALUES($1, $2)`
	db.QueryRow(query, u.Name, u.Age)

	return nil
}
