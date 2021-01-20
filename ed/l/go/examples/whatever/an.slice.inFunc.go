package main

import (
	"fmt"
)

func main() {
	// appendToSlice()
	appendToSlicePointer()
	// updateSlice()
}

func updateSlice() {
	s := make([]int, 1)
	sUpdate(s)
	fmt.Println("end:", s) // [5]
}

func sUpdate(s []int) {
	s[0] = 5
	fmt.Println("in sUpdate:", s) // [5]
}

func appendToSlice() {
	s := make([]int, 1)
	sAppend(s)
	fmt.Println("appendToSlice end:", s) // [0]
}

func sAppend(s []int) {
	s = append(s, 1)
	fmt.Println("in sAppend:", s) // [0 1]
}

func appendToSlicePointer() {
	s := make([]int, 1)
	sAppendToPointer(&s)
	fmt.Println("appendToSlicePointer end:", s) // [0 1]
}

func sAppendToPointer(s *[]int) {
	*s = append(*s, 1)
	fmt.Println("in sAppend:", s) // &[0 1]
}
