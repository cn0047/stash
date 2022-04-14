package main

import (
	"fmt"
)

// SomeExample - where S must be a slice type whose element type can be any type.
type SomeExample[S interface{ ~[]E }, E interface{}] []S

// SomeExample2 - same to SomeExample.
type SomeExample2[S ~[]E, E interface{}] []S

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

// AnyString - approximation constraint element.
type AnyString interface {
	~string
}

// SignedInteger - union constraint element.
type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Tree - simple tree.
type Tree[T interface{}] struct {
	left, right *Tree[T]
	value       T
}

func (t *Tree[T]) Lookup(x T) *Tree[T] {
	return nil
}

var stringTree Tree[string]

type structField interface {
	struct {
		a int
		x int
	} |
		struct {
			b int
			x float64
		} |
		struct {
			c int
			x uint64
		}
}

func main() {
	//simpleCase()
	runConvert()
}

func simpleCase() {
	var v Vector[int]
	v.Push(1)
	fmt.Printf("[simpleCase] %#v \n", v)
}

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func runConvert() {
	x := int64(10)
	r := convert[int32, int64](x)
	fmt.Printf("[runConvert] r=%+v \n", r)
}

func convert[To, From integer](from From) To {
	to := To(from)
	if From(to) != from {
		panic("convert fail: conversion out of range")
	}
	return to
}

//func Smallest[T constraints.Ordered](s []T) T {
//	r := s[0] // panics if slice is empty
//	for _, v := range s[1:] {
//		if v < r {
//			r = v
//		}
//	}
//	return r
//}
