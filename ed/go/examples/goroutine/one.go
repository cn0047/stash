package main

import (
	"time"
)

func main() {
	r := go f(1) // error
	println(r)
}

func f(x int) int {
	time.Sleep(100 * time.Millisecond)
	return x + 10
}
