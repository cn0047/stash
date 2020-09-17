package main

import (
	"fmt"
)

func f1() bool {
	fmt.Println("f1")
	return true
}

func f2() bool {
	fmt.Println("f2")
	return false
}

func main() {
	if f1() || f2() {
		fmt.Println("case 1")
	}
	fmt.Println("---")
	if f2() || f1() {
		fmt.Println("case 2")
	}
	fmt.Println("---")
	// if f1() | f2() { // invalid operation: f1() | f2() (operator | not defined on bool)
	// 	fmt.Println("case 3")
	// }
}

/*

f1
case 1
f2
f1
case 2

*/
