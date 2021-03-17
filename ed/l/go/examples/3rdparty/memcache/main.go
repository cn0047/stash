package main

import (
	"fmt"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

const (
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Black  = "\033[0m"
	Tick   = 500 * time.Millisecond
)

var (
	mc *memcache.Client
)

func main() {
	initMC()
	//f1()
	f4()
	fmt.Scanln()
}

func f4() {
	k := "k1"
	f2(k)
	for i := 0; i < 100; i++ {
		f3(k)
	}
}

func f3(k string) {
	go func() {
		for {
			v, _ := get(k)
			fmt.Printf("%s%s%s ", Black, v, Black)
			time.Sleep(Tick)
		}
	}()
}

func f2(k string) {
	set(k, "1", 0)
	go func() {
		for {
			v, _ := incr(k, 1)
			fmt.Printf("%s%v%s ", Purple, v, Black)
			time.Sleep(Tick)
		}
	}()
}

func f1() {
	k := "foo15"
	var ttl int32
	ttl = 30                                              // perfect
	ttl = int32(time.Now().Add(+20 * time.Second).Unix()) // imperfect

	ok, err := exists(k)
	fmt.Printf("[exists] 🎾 %#v \n", ok)
	if err != nil {
		fmt.Printf("🔴 %#v \n", err)
	}
	if ok {
		get(k)
	} else {
		set(k, "204", ttl)
		get(k)
	}
}

func initMC() {
	mc = memcache.New("127.0.0.1:11211")
}

func exists(key string) (bool, error) {
	_, err := mc.Get(key)
	if err != nil {
		if err == memcache.ErrCacheMiss {
			return false, nil
		}
		return false, fmt.Errorf("failed to get cache key, error: %v", err)
	}

	return true, nil
}

func incr(key string, delta uint64) (uint64, error) {
	val, err := mc.Increment(key, delta)
	return val, err
}

func set(key string, val string, ttl int32) error {
	err := mc.Set(&memcache.Item{Key: key, Value: []byte(val), Expiration: ttl})
	return err
}

func get(key string) ([]byte, error) {
	item, err := mc.Get(key)
	if err != nil {
		return nil, err
	}
	return item.Value, nil
}
