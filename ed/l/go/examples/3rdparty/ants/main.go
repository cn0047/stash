package main

import (
	"fmt"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
)

func main() {
	// submit()
	// poolInvoke()
	poolPanic()
}

func ping() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("ping")
}

func submit() {
	var wg sync.WaitGroup

	f := func() {
		ping()
		wg.Done()
	}

	n := 100
	for i := 0; i < n; i++ {
		wg.Add(1)
		_ = ants.Submit(f)
	}
	defer ants.Release()

	wg.Wait()
	fmt.Printf("goroutines count: %d\n", ants.Running())
	fmt.Printf("all tasks finished.\n")
}

var sum int32

func inc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("inc %d\n", n)
}

func incWithError(i interface{}) error {
	n := i.(int32)
	if n == int32(13) || n == int32(66) {
		return fmt.Errorf("incWithError error: bad number")
	}

	atomic.AddInt32(&sum, n)
	fmt.Printf("inc %d\n", n)

	return nil
}

func poolInvoke() {
	var wg sync.WaitGroup

	count := 10
	p, err := ants.NewPoolWithFunc(count, func(i interface{}) {
		inc(i)
		wg.Done()
	})
	if err != nil {
		// return fmt.Errorf("failed to create new ants pool, error: %v", err)
	}
	defer p.Release()

	n := 100
	for i := 0; i < n; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(1))
	}

	wg.Wait()
	fmt.Printf("goroutines count: %d\n", p.Running())
	fmt.Printf("sum = %d\n", sum)
}

func pf(p interface{}) {
	fmt.Printf("workerPool panic %$v", p)
	debug.PrintStack()
}

func poolPanic() {
	var wg sync.WaitGroup

	count := 10
	p, _ := ants.NewPool(count, ants.WithPanicHandler(pf))
	defer p.Release()

	n := 100
	for i := 0; i < n; i++ {
		wg.Add(1)
		err := p.Submit(func() {
			defer wg.Done()
			incWithError(int32(1))
		})
		if err != nil {
			fmt.Printf("failed submit func, error: %+v", err)
		}
	}

	wg.Wait()
	fmt.Printf("goroutines count: %d\n", p.Running())
	fmt.Printf("sum = %d\n", sum)
}
