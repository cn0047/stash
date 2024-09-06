// @see: https://leetcode.com/problems/validate-binary-search-tree
package main

import (
	"fmt"
	"math"
)

const (
	null = math.MaxInt - 1
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
	tree1 = []int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4} // false
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
	tree2 = []int{3, 5, 1, 6, 2, 0, 8, 7, 4, null, null, null, null, null, null, 10, null, null, 9} // false
	/*
	     5
	   /   \
	  1     4
	       / \
	       3  6
	*/
	tree3 = []int{5, 1, 4, null, null, 3, 6} // false
	/*
	     2
	   /   \
	  1     3
	*/
	tree4 = []int{2, 1, 3} // true
	/*
	     5
	   /   \
	  1     8
	       / \
	       7  9
	*/
	tree5 = []int{5, 1, 8, null, null, 7, 9} // true
	/*
	     2
	   /   \
	  10     3
	*/
	tree6 = []int{2, 10, 3} // false
	/*
	      5
	    /   \
	   4     6
	        / \
	       3   7
	*/
	tree7 = []int{5, 4, 6, null, null, 3, 7} // false
	tree8 = []int{2, 2, 2}                   // false
	/*
	      5
	    /   \
	   3     7
	        / \
	       4   8
	*/
	tree9 = []int{5, 3, 7, null, null, 4, 8} // false
	/*
	        5
	      /   \
	     3     8
	    / \   / \
	   1   4 6   9
	*/
	tree10 = []int{5, 3, 8, 1, 4, 6, 9} // true
	/*
			              120
		           ______|_______
			        /              \
			      70                140
			    /   \             /     \
		    50     100       130       160
		   /  \    /  \      /  \      /  \
			20  55  75  110  119  135  150  200
	*/
	tree11 = []int{120, 70, 140, 50, 100, 130, 160, 20, 55, 75, 110, 119, 135, 150, 200} // false
	tree12 = []int{-2147483648, null, 2147483647}                                        // true
)

func main() {
	var r bool
	r = isValidBST(toTree(tree1)) // false
	// r = isValidBST(toTree(tree2)) // false
	// r = isValidBST(toTree(tree3)) // false
	// r = isValidBST(toTree(tree4)) // true
	// r = isValidBST(toTree(tree5)) // true
	// r = isValidBST(toTree(tree6)) // false
	// r = isValidBST(toTree(tree7)) // false
	// r = isValidBST(toTree(tree8)) // false
	// r = isValidBST(toTree(tree9)) // false
	// r = isValidBST(toTree(tree10)) // true
	// r = isValidBST(toTree(tree11)) // false
	r = isValidBST(toTree(tree12)) // true
	fmt.Printf("res:%v\n", r)
}

func isValidBST(root *TreeNode) bool {
	return isValidSubBST(root.Left, math.MinInt, root.Val) && isValidSubBST(root.Right, root.Val, math.MaxInt)
}

func isValidSubBST(node *TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return isValidSubBST(node.Left, min, node.Val) && isValidSubBST(node.Right, node.Val, max)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// toTree creates tree out of slice.
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
