package main

import (
	"fmt"
)

const (
	CONST_VAL_A = 5 << iota
	CONST_VAL_B
	CONST_VAL_C
	CONST_VAL_D
)

func main() {
	fmt.Printf("%+v \n", CONST_VAL_A) // 5
	fmt.Printf("%+v \n", CONST_VAL_B) // 10
	fmt.Printf("%+v \n", CONST_VAL_C) // 20
	fmt.Printf("%+v \n", CONST_VAL_D) // 40
}
