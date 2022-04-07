package main

import (
	"fmt"
)

// Vector - simple vector.
type Vector[T any] []T

func (v *Vector[T]) Push(x T) {
	*v = append(*v, x)
}

// List - linked list.
type List[T any] struct {
	next *List[T]
	val  T
}

type P[T1, T2 any] struct {
	F *P[T2, T1] // INVALID; must be [T1, T2]
}

// AnyString - Approximation constraint element.
type AnyString interface {
	~string
}

// SignedInteger - union constraint element.
type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func main() {
	simpleCase()
}

func simpleCase() {
	var v Vector[int]
	v.Push(1)
	fmt.Printf("[simpleCase] %#v \n", v)
}
