package main

import (
	"time"
)

func main() {
	go func() {
		for {
			println(1)
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for {
			println(2)
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for {
			println(3)
			time.Sleep(1 * time.Second)
		}
	}()
	select {}
}

// in `ps aux` you'll see only 1 record
