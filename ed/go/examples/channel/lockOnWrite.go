package main

import (
	"fmt"
)

func main() {
	//one()
	two()
}

func two() {
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

func one() {
	c := make(chan int, 2)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("adding %v \n", i)
			c <- i
			fmt.Printf("added %v \n", i)
		}
	}()
	go func() {
	}()
	fmt.Scanln()
}
