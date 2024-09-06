package main

import (
	"fmt"
)

func main() {
	values := []int{}
	values = []int{50, 17, 76, 9, 23} // tree [ 50 17 76 9 23 _ _ _ _ _ _ ]
	//values = []int{76, 9, 23, 50, 17} // tree [ 23 9 76 _ 17 50 _ _ _ _ _ ]
	//values = []int{76, 50, 17, 9, 23} // tree [ 50 17 76 9 23 _ _ _ _ _ _ ]
	//values = []int{9, 19, 29, 5, 7, 39} // tree [ 19 9 29 5 _ _ 39 _ 7 _ _ _ _ ]

	t := &Tree{}
	for i := 0; i < len(values); i++ {
		fmt.Printf("\ninsert: %d", values[i])
		t.Insert(values[i])
		dump(t.Root)
	}
}

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	height int
}

func (n *Node) Height() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *Node) GetBalance() int {
	return n.Right.Height() - n.Left.Height()
}

func (n *Node) Insert(val int) *Node {
	if n == nil {
		return &Node{Value: val, height: 1}
	}
	if n.Value == val {
		return n
	}

	if val < n.Value {
		n.Left = n.Left.Insert(val)
	} else {
		n.Right = n.Right.Insert(val)
	}

	n.height = max(n.Left.Height(), n.Right.Height()) + 1

	return n
}

func (n *Node) rotateLeft() *Node {
	fmt.Printf("\nrotateLeft: %d", n.Value)
	r := n.Right
	n.Right = r.Left
	r.Left = n
	n.height = max(n.Left.Height(), n.Right.Height()) + 1
	r.height = max(r.Left.Height(), r.Right.Height()) + 1
	return r
}

func (n *Node) rotateRight() *Node {
	fmt.Printf("\nrotateRight: %d", n.Value)
	l := n.Left
	n.Left = l.Right
	l.Right = n
	n.height = max(n.Left.Height(), n.Right.Height()) + 1
	l.height = max(l.Left.Height(), l.Right.Height()) + 1
	return l
}

func (n *Node) rotateRightLeft() *Node {
	n.Right = n.Right.rotateRight()
	n = n.rotateLeft()
	n.height = max(n.Left.Height(), n.Right.Height()) + 1
	return n
}

func (n *Node) rotateLeftRight() *Node {
	n.Left = n.Left.rotateLeft()
	n = n.rotateRight()
	n.height = max(n.Left.Height(), n.Right.Height()) + 1
	return n
}

func (n *Node) rebalance() *Node {
	switch {
	case n.GetBalance() < -1 && n.Left.GetBalance() == -1:
		return n.rotateRight()
	case n.GetBalance() > 1 && n.Right.GetBalance() == 1:
		return n.rotateLeft()
	case n.GetBalance() < -1 && n.Left.GetBalance() == 1:
		return n.rotateLeftRight()
	case n.GetBalance() > 1 && n.Right.GetBalance() == -1:
		return n.rotateRightLeft()
	}

	return n
}

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(value int) {
	t.Root = t.Root.Insert(value)
	if t.Root.GetBalance() < -1 || t.Root.GetBalance() > 1 {
		t.rebalance()
	}
}

func (t *Tree) rebalance() {
	if t == nil || t.Root == nil {
		return
	}
	t.Root = t.Root.rebalance()
}

func dump(node *Node) {
	res := []*int{}
	q := []*Node{node}

	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if n == nil {
			res = append(res, nil)
		} else {
			q = append(q, n.Left, n.Right)
			res = append(res, &n.Value)
		}
	}

	fmt.Printf("\ntree [")
	for _, v := range res {
		if v == nil {
			fmt.Printf(" _")
		} else {
			fmt.Printf(" %d", *v)
		}
	}
	fmt.Printf(" ]")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
