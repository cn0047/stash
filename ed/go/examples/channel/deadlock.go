package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	ch <- 1 // fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
}
