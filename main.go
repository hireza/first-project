package main

import (
	handler "github.com/hireza/first-project/delivery/http"
	models "github.com/hireza/first-project/models"
	redisCount "github.com/hireza/first-project/redis"

	_ "github.com/lib/pq"
)

func main() {
	redisCount.InitialCount(redisCount.ConnectRedis())
	models.ConnectPosgres()
	handler.ConnectMux()
	// consumer.ConnectMSQ()
}
