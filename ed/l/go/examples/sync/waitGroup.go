package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//one()
	two()
}

func one() {
	fmt.Printf("[one] \n")
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

func two() {
	fmt.Printf("[two] \n")
	var wg sync.WaitGroup
	wg.Add(100) // !!!
	var x int32 = 0
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt32(&x, 1)
			wg.Done()
		}()
	}

	fmt.Println("Wait ...")
	wg.Wait()
	fmt.Println(atomic.LoadInt32(&x))
}

/*
[one]
Wait ...
99

[two]
Wait ...
100
*/
