package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	startedAt := time.Now().UnixNano()
	defer func() {
		finishedAt := time.Now().UnixNano()
		fmt.Printf("Took: %d microseconds \n\n", (finishedAt-startedAt)/1000)
	}()

	const numJobs = 1500
	// workerPerCPU(numJobs) // Took: 5688 microseconds
	goroutinePerJob(numJobs) // Took: 8175 microseconds
}

func goroutinePerJob(numJobs int) {
	jobs := make(chan int, numJobs) // ! fat chan
	results := make(chan int, numJobs)

	// produce jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	id := 0
	for j := range jobs {
		id++
		go func(id int, j int) {
			results <- do(id, j)
		}(id, j)
	}

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

func workerPerCPU(numJobs int) {
	workersCount := runtime.NumCPU()

	jobs := make(chan int, numJobs) // ! fat chan
	results := make(chan int, numJobs)

	for w := 1; w <= workersCount; w++ {
		go worker(w, jobs, results)
	}

	// produce jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for {
		select {
		case j, ok := <-jobs:
			if !ok {
				return
			}
			results <- do(id, j)
		}
	}
}

func do(id int, j int) int {
	fmt.Println("job", j, " worker", id)
	return 1000 + j
}
