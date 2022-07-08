package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	simpleExample()
}

// Result: RWMutex allows either one reader or one writer.
// read 1 	 0
// read 2 	 0
// read 2 	 0
// read 1 	 0
// write 	 0
// write 	 1
// write 	 2
// read 1 	 3
// read 2 	 3
// read 2 	 3
// read 1 	 3
// read 1 	 3
// read 2 	 3
func simpleExample() {
	var lock sync.RWMutex
	data := make(map[int]int)

	go func() {
		for i := 0; i < 100; i++ {
			lock.RLock()
			time.Sleep(time.Millisecond * 1000)
			fmt.Printf("read 1 \t %v \n", len(data))
			lock.RUnlock()
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			lock.RLock()
			time.Sleep(time.Millisecond * 1000)
			fmt.Printf("read 2 \t %v \n", len(data))
			lock.RUnlock()
		}
	}()

	time.Sleep(time.Second * 2)

	go func() {
		lock.Lock()
		for i := 0; i < 3; i++ {
			time.Sleep(time.Millisecond * 2000)
			fmt.Printf("write \t %v \n", i)
			data[i] = i
		}
		lock.Unlock()
	}()

	select {}
}
