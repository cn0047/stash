package main

import (
	"fmt"
)

func main() {
	next := f()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}

func f() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
