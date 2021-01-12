package main

import (
	"fmt"
	"sync"
	"time"
)

type MyMap struct {
	sync.Mutex
	m map[int]int
}

func (m *MyMap) set(i int) {
	m.Lock()
	defer m.Unlock()
	m.m[i] = i
}

func (m *MyMap) get(i int) int {
	m.Lock()
	defer m.Unlock()
	return m.m[i]
}

func main() {
	m := MyMap{m: make(map[int]int)}
	// Writes into map.
	go func() {
		for i := 0; i < 1000; i++ {
			m.set(i)
			time.Sleep(50 * time.Millisecond)
		}
	}()
	// Reads from map.
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(m.get(i))
			time.Sleep(50 * time.Millisecond)
		}
	}()
	fmt.Scanln()
}
