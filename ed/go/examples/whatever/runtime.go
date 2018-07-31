package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("NumCPU: %v\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %v\n", runtime.NumGoroutine())
	runtime.Gosched() // yields the processor, allowing other goroutines to runs
	runtime.Goexit() // terminates the goroutine
}
