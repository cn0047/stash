package main

import (
	"fmt"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

var (
	mc *memcache.Client
)

func main() {
	initMC()

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

func set(key string, val string, ttl int32) {
	err := mc.Set(&memcache.Item{Key: key, Value: []byte(val), Expiration: ttl})
	fmt.Printf("[SET] ✅ %#v \n", err)
}

func get(key string) {
	item, err := mc.Get(key)
	fmt.Printf("[get] 🔴 %#v \n", err)
	fmt.Printf("[get] 🎾 %#v \n \t val: %s \n", item, item.Value)
}
