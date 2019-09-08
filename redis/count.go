package redis

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

func ConnectRedis() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func InitialCount(c redis.Conn) error {
	_, err := c.Do("SET", "count", 1)
	if err != nil {
		log.Fatal(err)
	}

	get, err := redis.String(c.Do("GET", "count"))
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Count : ", get)

	return nil
}

func IncrementCount(c redis.Conn) (string, error) {
	_, err := c.Do("INCR", "count")
	if err != nil {
		log.Fatal(err)
	}

	get, err := redis.String(c.Do("GET", "count"))
	if err != nil {
		log.Fatal(err)
	}

	return get, nil
}
