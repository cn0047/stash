package main

import "fmt"

func main() {
	// zero()
	one()
	// two()
	three()
}

func zero() {
	ch := make(chan int, 0)
	ch <- 7
	v, ok := <-ch
	fmt.Printf("%+v, %+v \n", v, ok) // all goroutines are asleep - deadlock
}

func one() {
	ch := make(chan int, 1)
	ch <- 7
	v, ok := <-ch
	fmt.Printf("%+v, %+v \n", v, ok) // 7, true
}

func three() {
	ch := make(chan int, 1)
	close(ch)
	v, ok := <-ch
	fmt.Printf("%+v, %+v \n", v, ok) // 0, false
}

func two() {
	var ch chan int
	v, ok := <-ch
	fmt.Printf("%+v, %+v \n", v, ok) // all goroutines are asleep - deadlock
}
