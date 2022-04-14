package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Ordered interface {
	// "|" vertical bar here expresses a union of types,
	// "~string" - includes type string & all types declared with definitions like `type MyString string`,
	// where underlining type also string.
	constraints.Integer | constraints.Float | ~string
}

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func main() {
	f1()
}

func f1() {
	r := min[int](9, -1)
	fmt.Printf("[f1] r=%v \n", r) // [f1] r=-1
}

func scale[E constraints.Integer](s []E, c E) []E {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}
