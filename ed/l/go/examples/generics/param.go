package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
}

func (u User) String() string {
	return fmt.Sprintf("User: %s <%s>", u.Name, u.Email)
}

func main() {
	x := 1
	funcWithAnyTypeParam(x) // [funcWithAnyTypeParam] x value:1; x type: int

	sum(map[string]int64{"a": 1, "b": 2})  // [sum] m: map[string]int64{"a":1, "b":2}; res: 3
	sum2(map[string]int64{"a": 3, "b": 4}) // [sum2] m: map[string]int64{"a":3, "b":4}; res: 7

	reverse([]string{"1", "2", "3"}) // [reverse] s: []string{"3", "2", "1"}; res: []string{"3", "2", "1"}

	funcWith2AnyTypeParams("foo", "bar") // [funcWith2AnyTypeParams] x: "foo"; y: "bar"
	funcWith2AnyTypeParams(1, int64(2))  // [funcWith2AnyTypeParams] x: 1; y: 2
	//funcWith2AnyTypeParams("x", int64(2)) // cannot use "x" (untyped string constant) as int64 value in argument to funcWith2AnyTypeParams
}

func funcWithAnyTypeParam[myType any](x myType) {
	fmt.Printf("[funcWithAnyTypeParam] x value:%#v; x type: %T \n", x, x)
}

func funcWith2AnyTypeParams[T any](x T, y T) {
	fmt.Printf("[funcWith2AnyTypeParams] x: %#v; y: %#v \n", x, y)
}

func sum[K comparable, V int64 | float64](m map[K]V) V { // K,V - type parameters
	var s V
	for _, v := range m {
		s += v
	}
	fmt.Printf("[sum] m: %#v; res: %#v \n", m, s)
	return s
}

type Number interface {
	int64 | float64
}

func sum2[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	fmt.Printf("[sum2] m: %#v; res: %#v \n", m, s)
	return s
}

type Element interface {
	int | string
}

func reverse[E Element](s []E) {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
	fmt.Printf("[reverse] s: %#v; res: %#v \n", s, s)
}
