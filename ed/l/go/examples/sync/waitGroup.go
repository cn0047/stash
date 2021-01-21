package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var x int32 = 0
	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(1) // no guarantee it will be called before wg.Done()
			atomic.AddInt32(&x, 1)
			wg.Done() // no guarantee it will be called after wg.Add()
		}()
	}

	fmt.Println("Wait ...")
	wg.Wait()
	fmt.Println(atomic.LoadInt32(&x))
}

/*
Wait ...
99
*/
