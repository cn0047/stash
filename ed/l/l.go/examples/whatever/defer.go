package main

import (
	"fmt"
	"time"
)

func main() {
	three()
}

func three() {
	defer func(begin time.Time) {
		fmt.Printf("took %v \n", time.Since(begin))
	}(time.Now())

	fmt.Println("start")
	time.Sleep(2 * time.Second)
	fmt.Println("end")
}

func one() {
	defer println("top")
	defer println("bottom")

	println("main")
	/*
	   main
	   bottom
	   top
	*/
}

func two() {
	x := "foo"
	defer println(x)
	x = "bar"

	// Result:
	// foo
}
