package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(10)
	go f()
	println("Start:")
	time.Sleep(5000 * time.Millisecond)
}

func f() {
	for n := 1; n <= 20; n++ {
		time.Sleep(100 * time.Millisecond)
		println(n)
	}
}
