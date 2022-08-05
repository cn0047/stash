package main

import (
	"fmt"
)

func main() {
	root := initLinkedList1()
	r := dump(reverse(root))
	fmt.Printf("%v \n", r)
}

type Node struct {
	Val  int
	Next *Node
}

func initLinkedList1() *Node {
	var prev *Node
	var root *Node
	for i := 1; i <= 5; i++ {
		node := &Node{Val: i}
		if prev == nil {
			root = node
		} else {
			prev.Next = node
		}
		prev = node
	}

	return root
}

func reverse(head *Node) *Node {
	var prev *Node
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

func dump(head *Node) []int {
	r := make([]int, 0)
	for head != nil {
		r = append(r, head.Val)
		head = head.Next
	}

	return r
}
