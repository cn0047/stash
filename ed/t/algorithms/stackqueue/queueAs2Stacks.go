// see: https://www.hackerrank.com/challenges/ctci-queue-using-two-stacks/problem?h_l=interview&playlist_slugs%5B%5D%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D%5B%5D=stacks-queues
package main

import (
	"fmt"
)

func main() {
	s := NewStack()
	s.Enqueue("a")
	s.Enqueue("b")
	s.Enqueue("c")
	fmt.Println(GetBottomFromStack(s) == "a")
	s.Enqueue("d")
	s.Enqueue("e")
	fmt.Println(GetBottomFromStack(s) == "a")
	fmt.Println(s.Dequeue() == "e")
	DeleteFromStackBottom(s)
	fmt.Println(GetBottomFromStack(s))
}

func GetBottomFromStack(s *stack) string {
	tmp := NewStack()
	v := ""

	for !s.IsEmpty() {
		v = s.Dequeue()
		tmp.Enqueue(v)
	}
	for !tmp.IsEmpty() {
		s.Enqueue(tmp.Dequeue())
	}

	return v
}

func DeleteFromStackBottom(s *stack) string {
	tmp := NewStack()
	v := ""

	for {
		v = s.Dequeue()
		if s.IsEmpty() {
			break
		} else {
			tmp.Enqueue(v)
		}
	}
	for !tmp.IsEmpty() {
		s.Enqueue(tmp.Dequeue())
	}

	return v
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

func (s *stack) Get() string {
	v := s.Dequeue()
	s.Enqueue(v)

	return v
}
