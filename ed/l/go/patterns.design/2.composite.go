package main

import (
	"fmt"
)

type Composite interface {
	Do()
}

type Leaf struct {
	Value string
}

func (l Leaf) Do() {
	fmt.Println(l.Value)
}

type Node struct {
	Value string
	leafs []Composite
}

func (n *Node) Add(l Composite) { // pointer here is required for mutation leafs slice.
	n.leafs = append(n.leafs, l)
}

func (n Node) Do() {
	fmt.Printf("Node %s:\n", n.Value)
	for _, l := range n.leafs {
		l.Do()
	}
}

func main() {
	n := Node{Value: "root"}
	n.Add(Leaf{Value: "left"})
	n.Add(Leaf{Value: "right"})
	n2 := Node{Value: "middle-node"}
	n2.Add(Leaf{Value: "foo"})
	n2.Add(Leaf{Value: "bar"})
	n.Add(n2)
	n.Do()
}
