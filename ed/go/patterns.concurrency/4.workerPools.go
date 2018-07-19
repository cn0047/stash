package patterns_concurrency

import (
	"fmt"
	"time"
)

const (
	WorkersCount = 2
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= WorkersCount; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 7; j++ {
		jobs <- j
	}

	for a := 1; a <= 7; a++ {
		<-results
	}
	close(jobs)
}
