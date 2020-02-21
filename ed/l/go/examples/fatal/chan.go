package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil { // it won't recover fatal
			fmt.Printf("recovered: %v \n", r)
		}
	}()
	c()
}

func c() {
	c := make(chan int)
	c <- 1 // fatal error: all goroutines are asleep - deadlock!
}
