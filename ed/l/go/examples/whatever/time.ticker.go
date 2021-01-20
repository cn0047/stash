package main

import (
	"fmt"
	"time"
)

func f1() {
	threshold := 2 * time.Second
	ticker := time.NewTicker(threshold)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Printf("\n tick at: %v \n", t)
			default:
				fmt.Printf(".")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	fmt.Scanln()
}

func f2() {
	done := time.After(1 * time.Second)
	ticker := time.NewTicker(2 * time.Second)
	go func() {
		select {
		case <-done:
			print("done")
		case <-ticker.C:
			print("ticker")
		}
	}()

	time.Sleep(5 * time.Second)
}

func main() {
	f1()
	// f2()
}
