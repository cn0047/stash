package main

import (
	"fmt"
	"time"
)

func main() {
	// straightDefer()
	// valInDefer()
	poindterInDefer()
	// three()
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

func poindterInDefer() {
	x := "foo3"
	defer func(s *string) {
		println(*s)
	}(&x)
	x = "bar3"

	// Result:
	// bar3
}

func valInDefer() {
	x := "foo2"
	defer func(s string) {
		println(s)
	}(x)
	x = "bar2"

	// Result:
	// foo2
}

func straightDefer() {
	x := "foo"
	defer println(x)
	x = "bar"

	// Result:
	// foo
}
