// Binary Tree Level Order Traversal.
package main

import (
	"fmt"
)

const (
	null = 99999
)

func main() {
	// input := []int{3, 9, 20, null, null, 15, 7} // [[3] [9 20] [15 7]]
	// input := []int{3, 9, 20, 11, null, 15, 7} // [[3] [9 20] [11 15 7]]
	input := []int{1, 2, 3, 4, null, null, 5} // [[1] [2 3] [4 5]]
	t := toTree(input)
	r := levelOrder(t)
	fmt.Printf("===\n%v\n", r)
}

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrder returns slice of slices with values as "Binary Tree Level Order Traversal".
// @see: https://leetcode.com/problems/binary-tree-level-order-traversal/
func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return res
	}

	res = make([][]int, 0)
	// 2 queues to group ellements from one tree's level into one slice.
	queue := make([]*TreeNode, 0)     // for current tree's level.
	queueNext := make([]*TreeNode, 0) // for next tree's level.

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

// levelOrderBottom returns slice of slices with values as "Binary Tree Level Order Traversal"
// in order from bottom to top.
// @see: https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
func levelOrderBottom(root *TreeNode) (res [][]int) {
	if root == nil {
		return res
	}

	res = make([][]int, 0)
	// 2 queues to group ellements from one tree's level into one slice.
	queue := make([]*TreeNode, 0)     // for current tree's level.
	queueNext := make([]*TreeNode, 0) // for next tree's level.

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
			res = append([][]int{level}, res...)
		}
	}

	return res
}

// averageOfLevels returns average of Levels in Binary Tree.
// @see: https://leetcode.com/problems/average-of-levels-in-binary-tree/description/
func averageOfLevels(root *TreeNode) (res []float64) {
	if root == nil {
		return res
	}

	res = make([]float64, 0)
	// 2 queues to group ellements from one tree's level into one slice.
	queue := make([]*TreeNode, 0)     // for current tree's level.
	queueNext := make([]*TreeNode, 0) // for next tree's level.

	res = append(res, float64(root.Val))
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
			res = append(res, avg(level))
		}
	}

	return res
}

func avg(arr []int) float64 {
	c := len(arr)
	s := 0
	for _, v := range arr {
		s += v
	}

	return float64(s) / float64(c)
}

// averageOfLevels returns average of Levels in Binary Tree.
// @see: https://leetcode.com/problems/average-of-levels-in-binary-tree/description/
func averageOfLevels(root *TreeNode) (res []int) {
	if root == nil {
		return res
	}

	res = make([]int, 0)
	// 2 queues to group ellements from one tree's level into one slice.
	queue := make([]*TreeNode, 0)     // for current tree's level.
	queueNext := make([]*TreeNode, 0) // for next tree's level.

	res = append(res, root.Val)
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
			res = append(res, max(level))
		}
	}

	return res
}

func max(arr []int) int {
	m := arr[0]
	for _, v := range arr {
		if v > m {
			m = v
		}
	}

	return m
}
