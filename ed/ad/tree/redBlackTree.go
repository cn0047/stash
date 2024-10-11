package main

import (
	"fmt"
)

const (
	ColorBlack = 0
	ColorRed   = 1
)

type Node struct {
	val    int
	left   *Node
	right  *Node
	color  int
	parent *Node
}

type RedBlackTree struct {
	root *Node
	nil  *Node
}

func NewNode(val int, nilNode *Node) *Node {
	return &Node{
		val:    val,
		left:   nilNode,
		right:  nilNode,
		parent: nil,
		color:  ColorRed, // New node is red.
	}
}

func NewRedBlackTree() *RedBlackTree {
	n := &Node{color: ColorBlack}
	return &RedBlackTree{nil: n, root: n} // Root node is black.
}

func (r *RedBlackTree) Insert(data int) {
	newNode := NewNode(data, r.nil) // red.
	nilNode := r.nil
	n := r.root

	for n != r.nil {
		nilNode = n
		if newNode.val < n.val {
			n = n.left
		} else {
			n = n.right
		}
	}

	newNode.parent = nilNode

	if nilNode == r.nil {
		r.root = newNode
	} else if newNode.val < nilNode.val {
		nilNode.left = newNode
	} else {
		nilNode.right = newNode
	}

	if newNode.parent == r.nil {
		newNode.color = ColorBlack // Root node is black.
		return
	}

	if newNode.parent.parent == r.nil {
		return
	}

	r.fixInsert(newNode)
}

func (r *RedBlackTree) fixInsert(node *Node) {
	for node.parent.color == ColorRed {
		if node.parent == node.parent.parent.right {
			uncle := node.parent.parent.left
			if uncle.color == ColorRed {
				uncle.color = ColorBlack
				node.parent.color = ColorBlack
				node.parent.parent.color = ColorRed
				node = node.parent.parent
			} else {
				if node == node.parent.left {
					node = node.parent
					r.rightRotate(node)
				}
				node.parent.color = ColorBlack
				node.parent.parent.color = ColorRed
				r.leftRotate(node.parent.parent)
			}
		} else {
			uncle := node.parent.parent.right
			if uncle.color == ColorRed {
				uncle.color = ColorBlack
				node.parent.color = ColorBlack
				node.parent.parent.color = ColorRed
				node = node.parent.parent
			} else {
				if node == node.parent.right {
					node = node.parent
					r.leftRotate(node)
				}
				node.parent.color = ColorBlack
				node.parent.parent.color = ColorRed
				r.rightRotate(node.parent.parent)
			}
		}
		if node == r.root {
			break
		}
	}
	r.root.color = ColorBlack
}

func (r *RedBlackTree) leftRotate(node *Node) {
	n := node.right
	node.right = n.left
	if n.left != r.nil {
		n.left.parent = node
	}
	n.parent = node.parent
	if node.parent == r.nil {
		r.root = n
	} else if node == node.parent.left {
		node.parent.left = n
	} else {
		node.parent.right = n
	}
	n.left = node
	node.parent = n
}

func (r *RedBlackTree) rightRotate(node *Node) {
	n := node.left
	node.left = n.right
	if n.right != r.nil {
		n.right.parent = node
	}
	n.parent = node.parent
	if node.parent == r.nil {
		r.root = n
	} else if node == node.parent.right {
		node.parent.right = n
	} else {
		node.parent.left = n
	}
	n.right = node
	node.parent = n
}

func main() {
	values := []int{}
	values = []int{10, 20, 30, 15}                     // [20 10 30 0 15 0 0 _ _ 0 0 _ _ _ _ _ _ _ _]
	values = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}          // [4 2 6 1 3 5 8 0 0 0 0 0 0 7 9 _ _ _ _ _ _ _ _ _ _ _ _ 0 0 0 0 _ _ _ _ _ _ _ _]
	values = []int{21, 32, 60, 76, 100, 110, 145, 150} // [76 32 110 21 60 100 145 0 0 0 0 0 0 0 150 _ _ _ _ _ _ _ _ _ _ _ _ _ _ 0 0 _ _ _ _]
	values = []int{3, 5, 7, 8, 10, 15, 20, 30}         // [8 5 15 3 7 10 20 0 0 0 0 0 0 0 30 _ _ _ _ _ _ _ _ _ _ _ _ _ _ 0 0 _ _ _ _]
	t := NewRedBlackTree()
	for _, v := range values {
		t.Insert(v)
	}
	r := levelOrder(t.root)
	fmt.Printf("===\n%v", r)
}

// levelOrder returns slice which represents level order traversal for tree represented by provided root node.
func levelOrder(root *Node) (res []string) {
	queue := []*Node{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			res = append(res, "_")
		} else {
			res = append(res, fmt.Sprintf("%d", node.val))
			queue = append(queue, node.left, node.right)
		}
	}

	return res
}
