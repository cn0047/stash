package main

import (
	"time"
)

func main() {
	f()
}

func f() {
	go func() {
		panic(500)
	}()
	time.Sleep(time.Second * 2)
}
