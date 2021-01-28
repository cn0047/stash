package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	goredis "github.com/go-redis/redis"
	"github.com/go-redsync/redsync"
	redigo "github.com/gomodule/redigo/redis"
	radix "github.com/mediocregopher/radix/v3"
)

const (
	RedisAddr = ":6379" // xredis
)

func main() {
	//goRedisPing()
	//goRedisLock()
	//goRedisSet()
	//goRedisGet1()
	//goRedisGet2()

	//radixSet()
	//radixScript1()
	//radixCircularBuffer()
	radixCircularBuffer2()
}

func getRadixPool() *radix.Pool {
	f := radix.PoolConnFunc(func(network, addr string) (radix.Conn, error) {
		client, err := radix.Dial(network, addr)
		if err != nil {
			return nil, err
		}
		if err = client.Do(radix.Cmd(nil, "SELECT", "1")); err != nil {
			if e := client.Close(); e != nil {
				return nil, e
			}
			return nil, err
		}
		return client, nil
	})
	i := radix.PoolPingInterval(1 * time.Second)

	p, err := radix.NewPool("tcp", RedisAddr, 3, f, i)
	if err != nil {
		panic(fmt.Errorf("failed to create radix pool , error: %w", err))
	}

	return p
}

func radixSet() {
	p := getRadixPool()
	err := p.Do(radix.Cmd(nil, "SET", "rkey", "rval", "EX", "3600"))
	if err != nil {
		log.Printf("[radix set] error: %#v", err)
	}
}

func radixScript1() {
	s := `
		local k = redis.call('GET', KEYS[1])
		if k then
		  redis.call('EXPIRE', KEYS[1], ARGV[1])
		else
			return 'not found'
		end
		return k
	`
	res := ""
	script := radix.NewEvalScript(1, s).Cmd(&res, "rkey", "30")
	p := getRadixPool()
	err := p.Do(script)
	if err != nil {
		log.Printf("[radixScript set] error: %#v", err)
	}
	log.Printf("[radixScript set] res: %#v", res)
}

func radixCircularBuffer2() {
	p := getRadixPool()
	v := strconv.FormatInt(time.Now().Unix(), 10)
	s := `
		redis.call('LPUSH', KEYS[1], ARGV[1])
		redis.call('LTRIM', KEYS[1], 0, 5)
	`
	res := ""
	script := radix.NewEvalScript(1, s).Cmd(&res, "cb", v)
	err := p.Do(script)
	if err != nil {
		log.Printf("[radix circular buffer 2] error: %#v", err)
	}
	log.Printf("[radix circular buffer 2] res: %#v", res)
}

func radixCircularBuffer() {
	p := getRadixPool()
	v := strconv.FormatInt(time.Now().Unix(), 10)
	s := `
		redis.call('RPUSH', KEYS[1], ARGV[1])
		if (redis.call("LLEN", KEYS[1]) > 5) then redis.call("LPOP", KEYS[1]) end
	`
	res := ""
	script := radix.NewEvalScript(1, s).Cmd(&res, "cb", v)
	err := p.Do(script)
	if err != nil {
		log.Printf("[radix circular buffer] error: %#v", err)
	}
	log.Printf("[radix circular buffer] res: %#v", res)
}

func rediGoLock() {
	p := getRediGoPool(RedisAddr)
	rsc := redsync.New([]redsync.Pool{p})

	// lock
	m := rsc.NewMutex("user1", redsync.SetTries(3), redsync.SetExpiry(10*time.Second))
	err := m.Lock()
	if err != nil {
		fmt.Printf("[redigo] failed to acquire lock, error: %v \n", err)
		return
	}
	fmt.Println("[redigo] locked")

	// do some action
	time.Sleep(5 * time.Second)
	// unlock
	m.Unlock()
	fmt.Println("[redigo] unlocked")
}

func goRedisLock() {
	c := getGoRedisClient()

	// lock
	c.SetNX("user1", "1", 10*time.Second)
	// do some action
	time.Sleep(10 * time.Second)
	// unlock
	c.Del("user1")
}

func getRediGoPool(server string) *redigo.Pool {
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

func getRediGoClient() redigo.Conn {
	c, err := redigo.Dial("tcp", RedisAddr)
	if err != nil {
		panic(fmt.Errorf("[redigo] failed to get redigo connection , error: %w", err))
	}
	return c
}

func rediGoF() {
	c := getRediGoClient()
	_, err := c.Do("SET", "hello", "world")
	if err != nil {
		panic(fmt.Errorf("error3: %w", err))
	}
}

func getGoRedisClient3() *goredis.Client {
	opt := &goredis.ClusterOptions{
		Addrs:        []string{RedisAddr},
		MinIdleConns: 0,
		MaxConnAge:   0,
	}
	c := goredis.NewClusterClient(opt)
	_ = c
	return nil
}

func getGoRedisClient2() *goredis.Client {
	opt := &goredis.Options{
		Addr:         RedisAddr,
		MinIdleConns: 0,
		MaxConnAge:   0,
	}
	c := goredis.NewClient(opt)
	return c
}

func getGoRedisClient() *goredis.Client {
	c := goredis.NewClient(&goredis.Options{
		Network:  "tcp",
		Addr:     RedisAddr,
		Password: "",
		DB:       0, // default DB
	})
	return c
}

func goRedisPing() {
	c := getGoRedisClient()
	pong, err := c.Ping().Result()
	fmt.Println(pong, err)
}

func goRedisSet() {
	c := getGoRedisClient()
	r := c.Set("foo", "bar", 2*time.Second)
	log.Printf("[go-redis set] %+v", r)
}

func goRedisGet1() {
	c := getGoRedisClient()
	v := c.Exists("foo")
	e := v.Val() == int64(1)
	ex, eerr := v.Result()
	r := c.Get("foo")
	log.Printf("[go-redis get] exists: (%#v|%#v|%#v), %#v", e, ex, eerr, r.Val())
	// [get] exists: (true|1|<nil>), "bar"
}

func goRedisGet2() {
	time.Sleep(3 * time.Second)
	c := getGoRedisClient()
	v := c.Exists("foo")
	ex, eerr := v.Result()
	r := c.Get("foo")
	log.Printf("[go-redis get] exists: (%#v|%#v|%#v), %#v; %#v", v.Val(), ex, eerr, r.Val(), r.Err())
	// [get] exists: (0|0|<nil>), ""; "redis: nil"
}

func goRedisExists(key string) (bool, error) {
	c := getGoRedisClient()
	res, err := c.Exists(key).Result()
	if err != nil {
		return false, err
	}

	return res == int64(1), nil
}

var ErrorKeyAlreadyExists = errors.New("key already exists")

func goRedisSetNX(key, val string, ttl time.Duration) error {
	c := getGoRedisClient()
	success, err := c.SetNX(key, val, ttl).Result()
	if err != nil {
		return err
	}
	if !success {
		return ErrorKeyAlreadyExists
	}
	return nil
}
