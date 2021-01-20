package main

import (
	"fmt"
)

func main() {
	v, err := f4(0)
	if err != nil {
		fmt.Printf("Error in func main: %+v \n", err)
		return
	}
	fmt.Printf("Value: %+v \n", v)
}

func f4(n int) (v int, e error) {
	if n == 0 {
		e := fmt.Errorf("error")
		fmt.Printf("Error in func f: %+v \n", e)
		return
	}
	return
}

/*
ed/l/go/examples/traps/shadowingReturnError.go:20:3: e is shadowed during return
*/
