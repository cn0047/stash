package main

import (
	"fmt"
)

type VO struct {
	Value int
}

func main() {
	f1()
}

func f1() {
	v1 := &VO{Value: 1}
	v2 := *v1
	v3 := v1
	v2.Value = 2
	v3.Value = 3
	fmt.Printf("v1 = %+v \n", v1) // v1 = &{Value:3}
	fmt.Printf("v2 = %+v \n", v2) // v2 = {Value:2}
	fmt.Printf("v3 = %+v \n", v3) // v3 = &{Value:3}
}
