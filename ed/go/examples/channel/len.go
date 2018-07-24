package main

import (
	"fmt"
	"time"
)

func main() {
	n := 5
	c := make(chan int, n) // make(chan int, 0) equals to make(chan int)
	go f1(c, n)
	go f2(c)
	go f3(c)
	time.Sleep(time.Millisecond * 3000)
}

func f1(c chan int, n int) {
	for i := 0; i < n; i++ {
		c <- i
		fmt.Printf("added into chan value: %+v\n", i)
		time.Sleep(time.Millisecond * 10)
	}
}

func f2(c chan int) {
	for {
		v := len(c) // count of elements in chan
		fmt.Printf("len = %+v\n", v)
		time.Sleep(time.Millisecond * 100)
	}
}

func f3(c chan int) {
	for {
		v := <-c
		fmt.Printf("received value: %+v\n", v)
		time.Sleep(time.Millisecond * 200)
	}
}
