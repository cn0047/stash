// @see: https://leetcode.com/problems/same-tree
package main

import (
	"fmt"
	"math"
)

const (
	null = math.MaxInt - 1
)

func main() {
	tree1 := []int{2, 1, 3}
	tree2 := []int{2, 1, 3, 4}
	tree3 := []int{2, 3, 4, 1}

	var r bool
	r = isSameTree(toTree(tree1), toTree(tree2))
	r = isSameTree(toTree(tree1), toTree(tree1))
	r = isSameTree(toTree(tree2), toTree(tree3))
	fmt.Printf("===\n%v", r)
}

// isSameTree returns true in case both trees are same.
func isSameTree(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil || a.Val != b.Val {
		return false
	}

	return isSameTree(a.Left, b.Left) && isSameTree(a.Right, b.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
