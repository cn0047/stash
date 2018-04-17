package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan time.Time)
	go f(ch)
	fmt.Print("Start:")
	for {
		v := <-ch
		fmt.Printf("\rFrom channel: %v", v)
	}
}

func f(ch chan time.Time) {
	for {
		time.Sleep(100 * time.Millisecond)
		ch <- time.Now()
	}
}
