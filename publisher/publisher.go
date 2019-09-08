package publisher

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hireza/first-project/models"
	"github.com/nsqio/go-nsq"
)

func CreateUser(u *models.Users) {
	userJSON, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	config := nsq.NewConfig()
	p, err := nsq.NewProducer("127.0.0.1:4150", config)

	if err != nil {
		log.Panic(err)
	}

	err = p.Publish("tokped_users", []byte(string(userJSON)))
	if err != nil {
		log.Panic(err)
	}
}

func IncrementCount() {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer("127.0.0.1:4150", config)

	if err != nil {
		log.Panic(err)
	}

	err = p.Publish("tokped_count", []byte("a"))
	if err != nil {
		log.Panic(err)
	}
}
