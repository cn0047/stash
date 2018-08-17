package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go gr1(c)
	go gr2(c)
	time.Sleep(3 * time.Second)
}

func gr1(c chan int) {
	c <- 1
	c <- 2
	c <- 3
	close(c)
	//c <- 4 // panic: send on closed channel
}

func gr2(c chan int) {
	for v := range c {
		fmt.Printf("\nGot from channel: %+v", v)
	}
}

/*
Got from channel: 1
Got from channel: 2
Got from channel: 3
*/
