package main

import (
	"fmt"
	"time"
)

func main() {
	// one()
	// two()
	three()
}

func three() {
	c := make(chan int, 1)
	c <- 10
	fmt.Println("[3] ", <-c)
	close(c)
	c <- 11 // panic: send on closed channel
}

func two() {
	c := make(chan struct{})
	go func() {
		<-c
		fmt.Println("ok")
	}()
	time.Sleep(1 * time.Second)
	close(c) // after this will print ok
	fmt.Scanln()
}

func one() {
	c := make(chan struct{})
	go func() {
		<-c // hanging here
		fmt.Println("ok")
	}()
	fmt.Scanln()
}
