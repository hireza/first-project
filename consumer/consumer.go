package main

import (
	"log"
	"sync"

	"github.com/nsqio/go-nsq"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	decodeConfig := nsq.NewConfig()
	c, err := nsq.NewConsumer("tokped_users", "users", decodeConfig)
	if err != nil {
		log.Panic("Could not create consumer")
	}

	c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println("Success add User")
		log.Println(string(message.Body))
		return nil
	}))

	c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println("Counter +1")
		return nil
	}))

	err = c.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}

	log.Println("Ready to Consume....")
	wg.Wait()
}
