package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hireza/first-project/models"
	"github.com/nsqio/go-nsq"
)

func main() {
	user := &models.Users{
		ID:   1,
		Name: "reza",
		Age:  22,
	}

	userJSON, err := json.Marshal(user)
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
