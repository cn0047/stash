package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("NumCPU: %v\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %v\n", runtime.NumGoroutine()) // n of goroutines that currently exist
	fmt.Printf("Is windows: %v\n", runtime.GOOS == "windows")
	runtime.Gosched() // yields the processor, allowing other goroutines to runs
	runtime.Goexit()  // terminates the goroutine
}
