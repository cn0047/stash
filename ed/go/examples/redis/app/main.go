package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/go-redsync/redsync"
	redigo "github.com/gomodule/redigo/redis"
)

const (
	RedisAddr = ":6379" // xredis
)

func main() {
	lock1()
	return
	ping()
	set()
	get1()
	get2()
}

func lock2() {
	p := newRedigoPool(RedisAddr)
	rsc := redsync.New([]redsync.Pool{p})

	// lock
	m := rsc.NewMutex("user1", redsync.SetTries(3), redsync.SetExpiry(10*time.Second))
	err := m.Lock()
	if err != nil {
		fmt.Printf("failed to acquire lock, error: %v \n", err)
		return
	}
	fmt.Println("locked")

	// do some action
	time.Sleep(5*time.Second)
	// unlock
	m.Unlock()
	fmt.Println("unlocked")
}

func lock1() {
	c := getClient()

	// lock
	c.SetNX("user1", "1", 10*time.Second)
	// do some action
	time.Sleep(10*time.Second)
	// unlock
	c.Del("user1")
}

func newRedigoPool(server string) *redigo.Pool {
	p := redigo.Pool{
		MaxIdle:     3,
		IdleTimeout: 30 * time.Second,
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return &p
}

func getRedigoClient() redigo.Conn {
	c, err := redigo.Dial("tcp", RedisAddr)
	if err != nil {
		panic(fmt.Errorf("failed to get redigo connection , error: %w", err))
	}
	return c
}

func redigoF() {
	c := getRedigoClient()
	_, err := c.Do("SET", "hello", "world")
	if err != nil {
		panic(fmt.Errorf("error3: %w", err))
	}
}

func getClient3() *redis.Client {
	opt := &redis.ClusterOptions{
		Addrs:        []string{RedisAddr},
		MinIdleConns: 0,
		MaxConnAge:   0,
	}
	c := redis.NewClusterClient(opt)
	_ = c
	return nil
}

func getClient2() *redis.Client {
	opt := &redis.Options{
		Addr:         RedisAddr,
		MinIdleConns: 0,
		MaxConnAge:   0,
	}
	c := redis.NewClient(opt)
	return c
}

func getClient() *redis.Client {
	c := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     RedisAddr,
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
	time.Sleep(3 * time.Second)
	c := getClient()
	v := c.Exists("foo")
	ex, eerr := v.Result()
	r := c.Get("foo")
	log.Printf("[get] exists: (%#v|%#v|%#v), %#v; %#v", v.Val(), ex, eerr, r.Val(), r.Err())
	// [get] exists: (0|0|<nil>), ""; "redis: nil"
}
