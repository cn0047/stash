package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int)
	m[1] = 0
	m[2] = 0
	fmt.Println(m) // map[2:0 1:0]

	f1(m)
}

func f2(m map[int]int) {
	for k := 1; k <= 2; k++ {
		i := 10 + k
		m[i] = 0
	}
	fmt.Println(m) // map[2:0 11:0 12:0 1:0]
}

func f1(m map[int]int) {
	for k, _ := range m {
		fmt.Println("len:", len(m))
		i := 10 + k
		m[i] = 0
		fmt.Println(" add into map key: ", i)
	}
	fmt.Println(m) // map[12:0 21:0 1:0 2:0 11:0]
}