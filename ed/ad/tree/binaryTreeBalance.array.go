// @see: https://leetcode.com/problems/balance-a-binary-search-tree
package main

import (
	"fmt"
	"math"
)

const (
	null = math.MaxInt - 1
)

func main() {
	tree := []int{}
	// tree = []int{5, 1, 4, null, null, 3, 6}
	// tree = []int{1, null, 2, null, 3, null, 4, null, null}
	tree = []int{2, 1, 3}
	r := levelOrder(balanceBST(toTree(tree)))
	fmt.Printf("===\n%v", r)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// balanceBST performs BST balancing by serializing into slice and deserializing into structs.
func balanceBST(root *TreeNode) *TreeNode {
	arr := []int{}
	inOrder(root, &arr)
	fmt.Printf("inOrder: %v\n", arr)

	t := toTreeFromArray(arr)

	return t
}

func inOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrder(root.Right, arr)
}

func toTreeFromArray(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	mid := len(arr) / 2
	head := &TreeNode{Val: arr[mid]}
	head.Left = toTreeFromArray(arr[:mid])
	head.Right = toTreeFromArray(arr[mid+1:])

	return head
}

// @see: ed/ad/tree/binaryTreeValidation.go
func toTree(input []int) *TreeNode {
	if len(input) == 0 {
		return nil
	}

	root := &TreeNode{Val: input[0]}
	input = input[1:]
	queue := []*TreeNode{root}

	for len(input) > 0 {
		left := input[0]
		input = input[1:]
		right := null
		if len(input) > 0 {
			right = input[0]
			input = input[1:]
		}

		n := queue[0]
		queue = queue[1:]
		if left != null {
			n.Left = &TreeNode{Val: left}
			queue = append(queue, n.Left)
		}
		if right != null {
			n.Right = &TreeNode{Val: right}
			queue = append(queue, n.Right)
		}
	}

	return root
}

// @see: ed/ad/tree/levelOrder.go
func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return res
	}

	res = make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queueNext := make([]*TreeNode, 0)

	res = append(res, []int{root.Val})
	queueNext = append(queueNext, root)

	for len(queueNext) > 0 {
		queue = queueNext
		queueNext = make([]*TreeNode, 0)
		level := make([]int, 0)
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				level = append(level, node.Left.Val)
				queueNext = append(queueNext, node.Left)
			}
			if node.Right != nil {
				level = append(level, node.Right.Val)
				queueNext = append(queueNext, node.Right)
			}
		}
		if len(level) > 0 {
			res = append(res, level)
		}
	}

	return res
}
