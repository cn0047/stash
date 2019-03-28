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
	fmt.Printf("%+v", <-ch)
	fmt.Println("unblock")
}

/*
block
*/
