package main

import (
	"sync"
	"time"
)

var (
	once sync.Once
	t    time.Time
)

func main() {
	f3()
	f3()
	f3()
}

func i0() {
	t = time.Now()
}

func f3() {
	once.Do(i0)
	println(t.UnixNano())
}
