package main

import (
	"fmt"
)

type IntSet map[int]struct{}

func NewIntSet() IntSet {
	return make(map[int]struct{})
}

func (i IntSet) Put(v int) {
	i[v] = struct{}{}
}

func (i IntSet) ToSlice() []int {
	s := make([]int, 0, len(i))
	for v := range i {
		s = append(s, v)
	}
	return s
}

func case1() {
	set := NewIntSet()
	set.Put(5)
	set.Put(7)
	set.Put(7)
	set.Put(7)
	set.Put(7)
	set.Put(7)
	set.Put(9)
	fmt.Printf("[case1] set: %+v; slice: %+v", set, set.ToSlice())
}

func main() {
	case1()
}
