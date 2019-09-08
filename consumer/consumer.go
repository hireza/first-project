package main

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/hireza/first-project/models"
	redisCount "github.com/hireza/first-project/redis"
	"github.com/hireza/first-project/repository"
	"github.com/nsqio/go-nsq"

	_ "github.com/lib/pq"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	decodeConfig := nsq.NewConfig()
	c1, err := nsq.NewConsumer("tokped_users", "users", decodeConfig)
	if err != nil {
		log.Panic("Could not create consumer")
	}

	c2, err := nsq.NewConsumer("tokped_count", "count", decodeConfig)
	if err != nil {
		log.Panic("Could not create consumer")
	}

	c1.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		var u models.Users

		err = json.Unmarshal(message.Body, &u)
		if err != nil {
			log.Fatal(err)
		}

		if err := repository.CreateUser(models.ConnectPosgres(), u); err != nil {
			log.Fatal(err)
		}

		log.Println("Success add user : ", u)
		return nil
	}))

	c2.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		count, err := redisCount.IncrementCount(redisCount.ConnectRedis())
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Count : ", count)
		return nil
	}))

	err = c1.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}

	err = c2.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}

	log.Println("Ready to Consume....")
	wg.Wait()
}
