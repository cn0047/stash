package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int)
	m[1] = 0
	m[2] = 0
	fmt.Println(m) // map[2:0 1:0]

	rangeLoop(m) // ‼️ not deterministic
	forLoop(m)   // ok
}

func forLoop(m map[int]int) {
	for k := 1; k <= 2; k++ {
		i := 10 + k
		m[i] = 0
	}
	fmt.Println(m) // map[2:0 11:0 12:0 1:0]
}

// rangeLoop performs loop over map and inserts values into map.
func rangeLoop(m map[int]int) {
	for k, _ := range m {
		fmt.Println("len:", len(m))
		i := 10 + k
		m[i] = 0
		fmt.Println(" add into map key: ", i)
	}
	fmt.Println(m) // map[12:0 21:0 1:0 2:0 11:0]
}
