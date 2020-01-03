package main

import (
	"fmt"
)

func main() {
	case2() // ping
	case1() // fatal error: all goroutines are asleep - deadlock!
}

func case1() {
	messages := make(chan string)

	messages <- "ping"
	msg := <-messages
	fmt.Println(msg)
}

func case2() {
	messages := make(chan string, 1)

	messages <- "ping"
	msg := <-messages
	fmt.Println(msg)
}
