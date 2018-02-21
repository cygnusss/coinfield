package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var Client *redis.Client

// NewClient connects to Redis cache
func NewClient() {
	Client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, _ := Client.Ping().Result()

	fmt.Println(pong)
}
