package main

import "fmt"

func main() {
	one()
	// two()
	three()
}

func one() {
	ch := make(chan int, 1)
	ch <- 7
	v, ok := <-ch
	fmt.Printf("%+v, %+v \n", v, ok) // 7, true
}

func two() {
	var ch chan int
	v, ok := <-ch
	fmt.Printf("%+v, %+v \n", v, ok) // all goroutines are asleep - deadlock
}

func three() {
	ch := make(chan int, 1)
	close(ch)
	v, ok := <-ch
	fmt.Printf("%+v, %+v \n", v, ok) // o, false
}
