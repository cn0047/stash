package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	c := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     "xredis:6379",
		Password: "",
		DB:       0, // default DB
	})
	pong, err := c.Ping().Result()
	fmt.Println(pong, err)
}
