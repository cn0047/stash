package main

import (
	"fmt"
)

func main() {
	// write(nil)
	read(nil)
}

func write(m map[int]int) {
	m[1] = 1 // panic: assignment to entry in nil map
}

func read(m map[int]int) {
	fmt.Println(m[1]) // 0
	v, in := m[1]
	fmt.Println(v, in) // 0, false
}
