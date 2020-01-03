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
	if f2() || f1() {
		fmt.Println("case 2")
	}
}

/*

f1
case 1
f2
f1
case 2

*/
