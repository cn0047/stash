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

func (m *MyMap) get(i int) (int, bool) {
	m.Lock()
	defer m.Unlock()
	v, ok := m.m[i]
	return v, ok
}

func main() {
	m := MyMap{m: make(map[int]int)}
	// Writes into map.
	go func() {
		for i := 0; i < 100; i++ {
			m.set(i)
			time.Sleep(50 * time.Millisecond)
		}
	}()
	// Reads from map.
	go func() {
		for i := 0; i < 100; i++ {
			v, _ := m.get(i)
			fmt.Println(v)
			time.Sleep(50 * time.Millisecond)
		}
	}()
	fmt.Scanln()
}
