package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/hireza/first-project/models"
	"github.com/hireza/first-project/publisher"
	redisCount "github.com/hireza/first-project/redis"
	"github.com/hireza/first-project/repository"

	_ "github.com/lib/pq"
)

func ConnectMux() *mux.Router {
	Router := mux.NewRouter()
	Router.HandleFunc("/", GetUsers).Methods("GET")
	Router.HandleFunc("/", CreateUser).Methods("POST")
	log.Println("Listening on :1234...")
	http.ListenAndServe(":1234", Router)

	return Router
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := repository.GetUsers(models.ConnectPosgres())
	if err != nil {
		log.Fatal("Error fetching the data")
	}

	publisher.IncrementCount()

	c := redisCount.ConnectRedis()
	count, err := redis.String(c.Do("GET", "count"))
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(count + "\n"))
	json.NewEncoder(w).Encode(result)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u *models.Users
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&u); err != nil {
		log.Fatal("Invalid payload")
		return
	}

	defer r.Body.Close()

	publisher.CreateUser(u)
}
