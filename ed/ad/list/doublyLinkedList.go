package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Prev string
	Val  string
	Next string
}

type LinkedList map[string]*Node

func (l LinkedList) get(val string) *Node {
	node, ok := l[val]
	if !ok {
		return nil
	}

	return node
}

func (l LinkedList) set(val string, node *Node) {
	l[val] = node
}

func main() {
	//s := "a,b b,c d,c"
	//s := "a,b b,c c,d"
	//s := "a,b x,b"
	//s := "a,b a,c"
	s := "a,b b,c 1,2 2,3"
	buildLinkedList(s)
}

func buildLinkedList(s string) {
	l := make(LinkedList)

	pairs := strings.Split(s, " ")
	for _, pair := range pairs {
		values := strings.Split(pair, ",")
		valA := values[0]
		valB := values[1]
		nodeA := l.get(valA)
		nodeB := l.get(valB)

		if nodeA == nil && nodeB == nil {
			if len(l) > 0 {
				fmt.Printf("New chunk not connected to rest elements.\n")
				p(l)
			}
			l.set(valA, &Node{Val: valA, Next: valB})
			l.set(valB, &Node{Prev: valA, Val: valB})
		} else if nodeA == nil && nodeB != nil {
			if nodeB.Prev != "" {
				fmt.Printf("Node B already has prev value, pair: %s.\n", pair)
				p(l)
				return
			}
			nodeB.Prev = valA
			l.set(valB, nodeB)
			nodeA = &Node{Val: valA, Next: valB}
			l.set(valA, nodeA)
		} else if nodeA != nil && nodeB == nil {
			if nodeA.Next != "" {
				fmt.Printf("Node A already has next value, pair: %s.\n", pair)
				p(l)
				return
			}
			nodeA.Next = valB
			l.set(valA, nodeA)
			nodeB = &Node{Prev: valA, Val: valB}
			l.set(valB, nodeB)
		} else {
			fmt.Printf("Both nodes already have values.\n")
			p(l)
			return
		}
	}

	p(l)
}

func p(l LinkedList) {
	for _, n := range l {
		fmt.Printf("%v-%v-%v ", n.Prev, n.Val, n.Next)
	}
	fmt.Printf("\n")
}
