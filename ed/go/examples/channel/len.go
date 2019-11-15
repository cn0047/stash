package main

import (
	"fmt"
	"time"
)

func main() {
	// make(chan int) equals to make(chan int, 0) equals to make(chan int, 1)
	c := make(chan int, 2)

	n := 3
	go in(c, n)
	go printLength(c)
	go out(c)
	_, err := fmt.Scanln()
	if err != nil {
	}
}

func in(c chan int, n int) {
	for i := 0; i < n; i++ {
		c <- i
		fmt.Printf("added into chan value: %+v\n", i)
		time.Sleep(time.Millisecond * 10)
	}
}

func printLength(c chan int) {
	for {
		v := len(c) // count of elements in chan
		fmt.Printf("len = %+v\n", v)
		time.Sleep(time.Millisecond * 100)
	}
}

func out(c chan int) {
	for {
		v := <-c
		fmt.Printf("received value: %+v\n", v)
		time.Sleep(time.Millisecond * 200)
	}
}
