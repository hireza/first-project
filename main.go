package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hireza/first-project/models"
	"github.com/hireza/first-project/repository"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	DB     *sqlx.DB
	Router *mux.Router
}

type Users models.Users

func main() {
	conf := Config{}
	conf.connectPosgres()
	conf.connectMux()
}

func (r *Config) connectMux() {
	r.Router = mux.NewRouter()
	r.Router.HandleFunc("/", r.GetUsers).Methods("GET")
	r.Router.HandleFunc("/", r.CreateUser).Methods("POST")
	http.ListenAndServe(":1234", r.Router)
}

func (c *Config) connectPosgres() {
	var err error
	c.DB, err = sqlx.Open("postgres", "postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("Unable to connect : ", err)
	}

	log.Print("Connected to DB", c.DB)
}

func (c *Config) GetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := repository.GetUsers(c.DB)
	if err != nil {
		log.Fatal("Error fetching the data")
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (c *Config) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u Users
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&u); err != nil {
		log.Fatal("Invalid payload")
		return
	}

	defer r.Body.Close()
	if err := u.CreateUser(c.DB); err != nil {
		log.Fatal(err.error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}
