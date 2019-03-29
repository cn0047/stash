package main

import (
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	println(r.Intn(100))
	println(r.Intn(100))
	println(r.Intn(100))
}
