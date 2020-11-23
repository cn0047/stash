package main

import (
	"fmt"
)

func main() {
	// f1()
	f2()
}

func f1() {
	ch := make(chan int)
	ch <- 1 // fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
}

func f2() {
	go func() {
		ch := make(chan int)
		fmt.Println("before")
		ch <- 1 // hanging
		fmt.Println(<-ch)
	}()
	fmt.Scanln()
}
