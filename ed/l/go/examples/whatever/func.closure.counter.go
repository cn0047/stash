package main

import (
	"fmt"
)

func main() {
	next := f()
	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
	fmt.Println(next()) // 3
}

func f() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
