package main

import (
	"fmt"
)

func main() {
	s := NewStack()
	s.Enqueue("a")
	s.Enqueue("b")
	fmt.Println(s.Dequeue() == "b")
	fmt.Println(s.IsEmpty() == false)
	fmt.Println(s.Dequeue() == "a")
	fmt.Println(s.IsEmpty() == true)
}

type stack struct {
	values []string
}

func NewStack() *stack {
	return &stack{
		values: make([]string, 0),
	}
}

func (s *stack) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *stack) Enqueue(val string) {
	s.values = append(s.values, val)
}

func (s *stack) Dequeue() string {
	n := len(s.values) - 1
	val := s.values[n]
	s.values = s.values[:n]

	return val
}
