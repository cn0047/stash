package main

import (
	"fmt"
)

func main() {
	// simple()
	fallthroughFunc()
}

func simple() {
	i := 4
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5:
		fmt.Println("4 of 5")
	default:
		fmt.Println("oops...")
	}
}

func fallthroughFunc() {
	v := 42
	switch v {
	case 100:
		fmt.Println(100)
		fallthrough
	case 42:
		fmt.Println(42)
		fallthrough
	case 1:
		fmt.Println(1)
		fallthrough
	default:
		fmt.Println("default")
	}
	/*
		42
		1
		default
	*/
}
