package main

import (
	"fmt"
	"time"
)

var (
	queue = []int{}
)

func main() {
	go worker()
	for i := 1; i > 0; i++ {
		fmt.Println(lbrequest(i))
		time.Sleep(100 * time.Millisecond)
	}
}

// lbrequest adds request to queue (buckets) if it has space (only 3 items in queue allowed).
func lbrequest(n int) string {
	if len(queue) == 3 {
		return fmt.Sprintf("%d \t sorry", n)
	}
	queue = append(queue, n)
	return fmt.Sprintf("%d \t enqueued", n)
}

// worker controls the pace (fulfilled requests per time range).
func worker() {
	for {
		if len(queue) == 0 {
			time.Sleep(10 * time.Millisecond)
			continue
		}
		n := queue[0]
		queue = queue[1:]
		time.Sleep(1 * time.Second)
		fmt.Printf("%d \t fulfilled\n", n)
	}
}
