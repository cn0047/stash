// @see: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree
package main

import (
	"fmt"
	"math"
)

const (
	null = 9999999
)

var (
	/*
		     3
		   /   \
		  5     1
		 / \   / \
		6   2  0  8
		   / \
		  7   4
	*/
	tree1 = []int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4}

	/*
	           3
	         /   \
	        5     1
	       / \   / \
	      6   2  0  8
	     / \
	    7   4
	   /     \
	  10      9
	*/
	tree2 = []int{3, 5, 1, 6, 2, 0, 8, 7, 4, null, null, null, null, null, null, 10, null, null, 9}
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}

func main() {
	r := lowestCommonAncestor(
		// toTree(tree1), &TreeNode{Val: 5}, &TreeNode{Val: 1}, // 3
		// toTree(tree1), &TreeNode{Val: 5}, &TreeNode{Val: 4}, // 5
		// toTree(tree1), &TreeNode{Val: 6}, &TreeNode{Val: 4}, // 5
		// toTree(tree1), &TreeNode{Val: 6}, &TreeNode{Val: 8}, // 3
		// toTree(tree1), &TreeNode{Val: 7}, &TreeNode{Val: 8}, // 3
		// toTree(tree1), &TreeNode{Val: 8}, &TreeNode{Val: 7}, // 3
		// toTree(tree1), &TreeNode{Val: 7}, &TreeNode{Val: 4}, // 2
		// toTree(tree1), &TreeNode{Val: 6}, &TreeNode{Val: 88}, // 6
		// toTree(tree2), &TreeNode{Val: 7}, &TreeNode{Val: 0}, // 3
		// toTree(tree2), &TreeNode{Val: 10}, &TreeNode{Val: 9}, // 6
		toTree(tree2), &TreeNode{Val: 8}, &TreeNode{Val: 9}, // 3
	)
	fmt.Printf("Result:\n%v\n", r.Val)

	// t, m, h := toTreeWithInfo([]int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4})
	// printTree(t, m, h)
	// fmt.Printf("Result:\n%v %v\n", m, h)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func toTreeWithInfo(input []int) (rootNode *TreeNode, maxVal int, height int) {
	if len(input) == 0 {
		return nil, 0, 0
	}

	h := int(math.Floor(math.Log2(float64(len(input) + 1)))) // floor(log2(n + 1))

	root := &TreeNode{Val: input[0]}
	max := input[0]
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
		if left > max && left != null {
			max = left
		}
		if right > max && right != null {
			max = right
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

	return root, max, h
}

func printTree(root *TreeNode, max int, height int) {
	fmt.Printf("Result:\n%v\n", root)
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
