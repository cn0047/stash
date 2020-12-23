package main

import (
	"fmt"
)

func main() {
	// avoidHangWithSelect()
	hang()
}

func avoidHangWithSelect() {
	c := make(chan int, 2)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("adding %v \n", i)
			select {
			case c <- i:
				fmt.Printf("added %v \n", i)
			default:
				fmt.Printf("did not add %v \n", i)
			}
		}
	}()
	go func() {
	}()
	fmt.Scanln()
}

/*
adding 0
added 0
adding 1
added 1
adding 2
did not add 2
adding 3
did not add 3
adding 4
did not add 4
*/

func hang() {
	c := make(chan int, 2)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("adding \t %v \n", i)
			c <- i // hanging for 3th iteration
			fmt.Printf("added \t %v \n", i)
		}
	}()
	go func() {
	}()
	fmt.Scanln()
}
