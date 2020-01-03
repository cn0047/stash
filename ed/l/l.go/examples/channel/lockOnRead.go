package main

import (
	"fmt"
)

func main() {
	//one()
	//two()
	three()
}

func three() {
	c := make(chan int, 3)
	for i := 0; i < 3; i++ {
		c <- i
	}
	go func() {
		for i := 0; i < 4; i++ {
			fmt.Printf("reading\n")
			select {
			case v := <-c:
				fmt.Printf("got %v \n", v)
			default:
				fmt.Printf("did not read \n")
			}
		}
		fmt.Printf("exit \n")
	}()
	go func() {
	}()
	fmt.Scanln()
}

func two() {
	c := make(chan int, 3)
	for i := 0; i < 3; i++ {
		c <- i
	}
	go func() {
		for i := 0; i < 4; i++ {
			fmt.Printf("reading\n")
			v, ok := <-c
			fmt.Printf("got %v, ok %v \n", v, ok)
		}
	}()
	go func() {
	}()
	fmt.Scanln()
}

func one() {
	c := make(chan int, 3)
	for i := 0; i < 3; i++ {
		c <- i
	}
	go func() {
		for i := 0; i < 4; i++ {
			fmt.Printf("reading\n")
			v := <-c
			fmt.Printf("got %v \n", v)
		}
	}()
	go func() {
	}()
	fmt.Scanln()
}
