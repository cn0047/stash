package main

import (
	"errors"
	"fmt"
)

var (
	ERR1 = errors.New("ERROR_#1")
	ERR2 = errors.New("ERROR_#2")
)

func main() {
	r := f1()
	fmt.Printf("Main result: %v. \n", r)
}

func f1() int {
	defer catch(ERR2, func() {
		fmt.Printf("Caught: %v. \n", ERR2)
		//return -1 // won't work
	})
	f2()
	return 1
}

func f2() {
	panic(ERR2)
}

func catch(e error, cb func()) {
	r := recover()

	if r == e {
		cb()
		return
	}

	if e != nil {
		panic(r)
	}
}
