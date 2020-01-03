package main

import (
	"fmt"
)

func worker(id int, jobs <-chan rune, results chan<- string) {
	for j := range jobs {
		fmt.Println("🎾 worker", id, "busy with job:", string(j))
		results <- fmt.Sprintf("done job: %s", string(j))
	}
	fmt.Println("🔴 worker", id, "exited")
}

func main() {
	data := []rune{} //{'a', 'b', 'c', 'd', 'e'}

	n := len(data)
	jobs := make(chan rune, n)
	results := make(chan string, n)

	workersCount := 13
	for w := 0; w < workersCount; w++ {
		go worker(w, jobs, results)
	}

	for _, v := range data {
		jobs <- v
	}
	close(jobs)

	for i := 0; i < n; i++ {
		fmt.Println(<-results)
	}
}
