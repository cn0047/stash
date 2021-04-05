package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m     map[int]int
	mutex *sync.Mutex
)

func main() {
	f1()
}

func f1() {
	m = make(map[int]int, 10)
	mutex = &sync.Mutex{}

	go write()
	go read()

	fmt.Scanln()
}

func write() {
	for i := 0; i < 10; i++ {
		mutex.Lock()
		m[i] = i
		mutex.Unlock()

		time.Sleep(90 * time.Millisecond)
	}
}

func read() {
	for i := 0; i < 10; i++ {
		mutex.Lock()
		fmt.Printf("%+v\n", m)
		mutex.Unlock()

		time.Sleep(90 * time.Millisecond)
	}
}
