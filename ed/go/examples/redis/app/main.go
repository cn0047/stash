package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

func main() {
	ping()
	set()
	get1()
	get2()
}

func getClient() (*redis.Client) {
	c := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     "xredis:6379",
		Password: "",
		DB:       0, // default DB
	})
	return c
}

func ping() {
	c := getClient()
	pong, err := c.Ping().Result()
	fmt.Println(pong, err)
}

func set() {
	c := getClient()
	r := c.Set("foo", "bar", 2*time.Second)
	log.Printf("[set] %+v", r)
}

func get1() {
	c := getClient()
	v := c.Exists("foo")
	e := v.Val() == int64(1)
	ex, eerr := v.Result()
	r := c.Get("foo")
	log.Printf("[get] exists: (%#v|%#v|%#v), %#v", e, ex, eerr, r.Val())
	// [get] exists: (true|1|<nil>), "bar"
}

func get2() {
	time.Sleep(3*time.Second)
	c := getClient()
	v := c.Exists("foo")
	ex, eerr := v.Result()
	r := c.Get("foo")
	log.Printf("[get] exists: (%#v|%#v|%#v), %#v; %#v", v.Val(), ex, eerr, r.Val(), r.Err())
	// [get] exists: (0|0|<nil>), ""; "redis: nil"
}
