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

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// @see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
