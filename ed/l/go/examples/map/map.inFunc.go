package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int)
	f(m)
	fmt.Println("end:", m) // end: map[2:2]
}

func f(m map[int]int) {
	m[2] = 2
	fmt.Println("in f:", m) // in f: map[2:2]
}
