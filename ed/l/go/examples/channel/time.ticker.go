package main

import (
	"fmt"
	"time"
)

func main() {
	// f1()
	simpleTicker()
	// f3()
}

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

func simpleTicker() {
	done := time.After(7 * time.Second)
	ticker := time.NewTicker(2 * time.Second)
	go func() {
		defer fmt.Println("exit")
		for {
			select {
			case <-done:
				fmt.Println("done")
				ticker.Stop()
				return
			case <-ticker.C:
				fmt.Println("ticker")
			}
		}
	}()

	fmt.Scanln()
}

func f3() {
	timer := time.NewTimer(time.Second)
	defer timer.Stop()

	messages := make(chan int)
	go func() {
		for i := 0; i < 1e5; i++ {
			time.Sleep(500 * time.Millisecond)
			messages <- i
		}
		close(messages)
	}()

	for {
		select {
		case <-timer.C:
			fmt.Println("Timeout")
			return
		case msg := <-messages:
			fmt.Println(msg)
			// Important.
			if !timer.Stop() {
				v := <-timer.C
				fmt.Printf("timer.C=%v", v)
			}
		default:
			time.Sleep(100 * time.Millisecond)
			fmt.Print(".")
		}

		// Reset to reuse.
		// timer.Reset(time.Second)
	}
}
