package main

import "fmt"

type Node struct {
	Tree  *BinaryTree
	Deep  int32 // height
	Value int32
	Root  int32
	Left  int32
	Right int32
}

func (n *Node) GetNode(index int32) *Node {
	node, ok := n.Tree.Tree[index]
	if !ok {
		return nil
	}

	return node
}

func (n *Node) GetLeft() *Node {
	return n.GetNode(n.Left)
}

func (n *Node) GetRoot() *Node {
	return n.GetNode(n.Root)
}

func (n *Node) GetRight() *Node {
	return n.GetNode(n.Right)
}

func (n *Node) GetValuesInOrderTraversal() []int32 {
	result := make([]int32, 0)

	l := n.GetLeft()
	if l != nil {
		result = append([]int32{}, l.GetValuesInOrderTraversal()...)
	}

	result = append(result, n.Value)

	r := n.GetRight()
	if r != nil {
		result = append(result, r.GetValuesInOrderTraversal()...)
	}

	return result
}

// SwapChildrenAtDeep performs swapping of node's children at certain level,
// level is deep or height of tree,
// where node's deep value must be: n.Deep % deepValue == 0.
func (n *Node) SwapChildrenAtDeep(deepValue int32) {
	if n.Deep%deepValue == 0 {
		tmp := n.Left
		n.Left = n.Right
		n.Right = tmp
	}

	l := n.GetLeft()
	if l != nil {
		l.SwapChildrenAtDeep(deepValue)
	}

	r := n.GetRight()
	if r != nil {
		r.SwapChildrenAtDeep(deepValue)
	}
}

func (n *Node) SwapChildren() {
	l := n.GetLeft()
	if l != nil {
		l.SwapChildren()
	}

	r := n.GetRight()
	if r != nil {
		r.SwapChildren()
	}

	tmp := n.Left
	n.Left = n.Right
	n.Right = tmp
}

type BinaryTree struct {
	Tree map[int32]*Node
	Root int32
}

func (b *BinaryTree) AddNode(value int32, left int32, right int32) {
	// Init map for tree data.
	if b.Tree == nil {
		b.Tree = make(map[int32]*Node)
	}

	node, ok := b.Tree[value]
	if !ok {
		b.Root = value
		node = &Node{Tree: b, Deep: 1, Value: value}
	}
	node.Left = left
	node.Right = right
	b.Tree[value] = node

	_, lOk := b.Tree[left]
	if left != -1 && !lOk {
		b.Tree[left] = &Node{Tree: b, Deep: node.Deep + 1, Value: left, Root: value}
	}

	_, rOk := b.Tree[right]
	if right != -1 && !rOk {
		b.Tree[right] = &Node{Tree: b, Deep: node.Deep + 1, Value: right, Root: value}
	}
}

func (b *BinaryTree) GetRoot() *Node {
	node, ok := b.Tree[b.Root]
	if !ok {
		return nil
	}

	return node
}

func main() {
	i := [][]int32{
		//{2, 3},
		//{-1, -1},
		//{-1, -1},
		//
		//{2, 3},
		//{-1, 4},
		//{-1, 5},
		//{-1, -1},
		//{-1, -1},
		//
		{2, 3},
		{4, -1},
		{5, -1},
		{6, -1},
		{7, 8},
		{-1, 9},
		{-1, -1},
		{10, 11},
		{-1, -1},
		{-1, -1},
		{-1, -1},
	}
	q := []int32{
		//1,
		//1,
		//
		//2,
		//
		2,
		4,
	}
	r := swapNodes(i, q)
	fmt.Println(r)
}

func swapNodes(indexes [][]int32, queries []int32) (result [][]int32) {
	tree := BinaryTree{}
	for i, ind := range indexes {
		tree.AddNode(int32(i+1), ind[0], ind[1])
	}

	for _, k := range queries {
		tree.GetRoot().SwapChildrenAtDeep(k)
		result = append(result, tree.Tree[1].GetValuesInOrderTraversal())
	}

	return result
}
