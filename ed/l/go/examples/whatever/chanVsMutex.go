package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startedAt := time.Now().UnixNano()
	defer func() {
		finishedAt := time.Now().UnixNano()
		fmt.Printf("Took: %d microseconds \n\n", (finishedAt-startedAt)/1000)
	}()

	n := 100
	// withChan(n) // Took: 258 microseconds
	mithMutex(n) // Took: 395 microseconds
}

func withChan(n int) {
	errs := make(chan error, n)

	wg := &sync.WaitGroup{}
	wg.Add(n)
	for w := 1; w <= n; w++ {
		go func(w int) {
			errs <- fmt.Errorf("err: %v", w)
			wg.Done()
		}(w)
	}
	wg.Wait()
	close(errs)

	es := make([]error, 0, n)
	for e := range errs {
		es = append(es, e)
	}
	fmt.Println("len:", len(es))
}

func mithMutex(n int) {
	es := make([]error, 0, n)

	wg := &sync.WaitGroup{}
	wg.Add(n)
	m := sync.Mutex{}
	for w := 1; w <= n; w++ {
		go func(w int) {
			m.Lock()
			e := fmt.Errorf("err: %v", n)
			es = append(es, e)
			m.Unlock()
			wg.Done()
		}(w)
	}
	wg.Wait()

	fmt.Println("len:", len(es))
}
