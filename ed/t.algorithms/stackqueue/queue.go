package main

import (
	"fmt"
)

func main() {
	q := NewQueue()
	q.Enqueue("a")
	q.Enqueue("b")
	fmt.Println(q.Dequeue() == "a")
	fmt.Println(q.IsEmpty() == false)
	fmt.Println(q.Dequeue() == "b")
	fmt.Println(q.IsEmpty() == true)
}

type queue struct {
	values []string
}

func NewQueue() *queue {
	return &queue{
		values: make([]string, 0),
	}
}

func (q *queue) IsEmpty() bool {
	return len(q.values) == 0
}

func (q *queue) Enqueue(val string) {
	q.values = append(q.values, val)
}

func (q *queue) Dequeue() string {
	val := q.values[0]
	q.values = q.values[1:]

	return val
}
