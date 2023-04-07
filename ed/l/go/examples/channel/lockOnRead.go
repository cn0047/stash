package main

import (
	"fmt"
)

func main() {
	// hangToRead()
	// hangToReadWithCommaOk()
	avoidHangWithSelect()
}

func avoidHangWithSelect() {
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
				fmt.Printf("got nothing \n")
			}
		}
		fmt.Printf("exit \n")
	}()
	go func() {
	}()
	fmt.Scanln()
}
/*
Result:
reading
got 0
reading
got 1
reading
got 2
reading
got nothing
exit
*/

func hangToReadWithCommaOk() {
	c := make(chan int, 3)
	for i := 0; i < 3; i++ {
		c <- i
	}
	go func() {
		for i := 0; i < 4; i++ {
			fmt.Printf("reading\n")
			v, ok := <-c // hanging for 4th iteration
			fmt.Printf("got %v, ok %v \n", v, ok)
		}
	}()
	go func() {
	}()
	fmt.Scanln()
}

func hangToRead() {
	c := make(chan int, 3)
	for i := 0; i < 3; i++ {
		c <- i
	}
	go func() {
		for i := 0; i < 4; i++ {
			fmt.Printf("reading\n")
			v := <-c // hanging for 4th iteration
			fmt.Printf("got %v \n", v)
		}
	}()
	go func() {
	}()
	fmt.Scanln()
}
