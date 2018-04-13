package main

import (
	"math/rand"
	"time"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 10; i++ {
		randomNum := random(1, 2000)
		println(randomNum)
	}
}
