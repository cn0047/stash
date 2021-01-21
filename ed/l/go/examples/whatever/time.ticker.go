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
		//timer.Reset(time.Second)
	}
}

func main() {
	// f1()
	// f2()
	f3()
}
