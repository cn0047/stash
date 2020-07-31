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
	v2.Value = 2
	fmt.Printf("v1 = %+v \n", v1)
	fmt.Printf("v2 = %+v \n", v2)
}
