package main

import (
	"fmt"
)

func main() {
	go f(nil)
	fmt.Scanln()
}

func f(ch chan int) {
	fmt.Println("block")
	ch <- 1
	fmt.Println("unblock")
}

/*
block
*/
