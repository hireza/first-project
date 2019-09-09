package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/hireza/first-project/models"
	"github.com/hireza/first-project/publisher"
	"github.com/hireza/first-project/repository"

	redisCount "github.com/hireza/first-project/redis"

	_ "github.com/lib/pq"
)

type home struct {
	Users []models.Users `json:"users"`
	Count string         `json:"count"`
}

func setHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//anyone can make a CORS request (not recommended in production)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//only allow GET, POST, and OPTIONS
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		//Since I was building a REST API that returned JSON, I set the content type to JSON here.
		w.Header().Set("Content-Type", "application/json")
		//Allow requests to have the following headers
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, cache-control")
		//if it's just an OPTIONS request, nothing other than the headers in the response is needed.
		//This is essential because you don't need to handle the OPTIONS requests in your handlers now
		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func ConnectMux() *mux.Router {
	// corsObj := handlers.AllowedOrigins([]string{"*"})
	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	// originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	Router := mux.NewRouter()
	Router.HandleFunc("/", GetUsers).Methods("GET")
	Router.HandleFunc("/", CreateUser).Methods("POST")
	log.Println("Listening on :1234...")
	http.ListenAndServe(":1234", setHeaders(Router))

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

	oke1 := &home{
		Users: result,
		Count: count,
	}

	oke2, err := json.Marshal(oke1)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	json.NewEncoder(w).Encode(string(oke2))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u *models.Users
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&u); err != nil {
		log.Fatal("Invalid payload")
		return
	}

	log.Println(u)

	defer r.Body.Close()

	publisher.CreateUser(u)
}
