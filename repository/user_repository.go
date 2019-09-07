package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/hireza/first-project/models"
	"github.com/jmoiron/sqlx"
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

func (u *Users) CreateUser(db *sql.DB) error {
	query := "INSERT INTO user_collection(Name, Age) VALUES($1, $2) RETURNING ID"
	err := db.QueryRow(query, u.Name, u.Age).Scan(u.ID)

	if err != nil {
		return err
	}

	return nil
}

func (u *mysqlUser) GetUser(db *sql.DB) error {
	return errors.New("Not implemented")
}
