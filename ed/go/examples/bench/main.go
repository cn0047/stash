package main

import (
	"fmt"
	"time"

	"./fibonacci"
)

func main() {
	go spinner(5 * time.Millisecond)
	const n = 40
	fibN := fibonacci.Fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
