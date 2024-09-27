// @see: https://leetcode.com/problems/binary-tree-maximum-path-sum
package main

import (
	"fmt"
	"math"
)

func main() {
	r01 := &TreeNode{Val: 1}
	r01.Left = &TreeNode{Val: 2}
	r01.Right = &TreeNode{Val: 3}

	r02 := &TreeNode{Val: 10}
	r02.Left = &TreeNode{Val: 5}
	r02.Right = &TreeNode{Val: 20}
	r02.Right.Left = &TreeNode{Val: 15}
	r02.Right.Right = &TreeNode{Val: 25}

	r03 := &TreeNode{Val: -10}
	r03.Left = &TreeNode{Val: 9}
	r03.Right = &TreeNode{Val: 20}
	r03.Right.Left = &TreeNode{Val: 15}
	r03.Right.Right = &TreeNode{Val: 7}

	s := &Solution{}
	res := 0
	res = s.maxPathSum(r01)
	res = s.maxPathSum(r02)
	res = s.maxPathSum(r03)
	fmt.Printf("===\n%v\n", res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Solution struct {
	answer int
}

func (s *Solution) dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := s.dfs(root.Left)
	r := s.dfs(root.Right)
	s.answer = max(s.answer, l+r+root.Val)
	res := max(0, root.Val+max(l, r))

	return res
}

func (s *Solution) maxPathSum(root *TreeNode) int {
	s.answer = math.MinInt32
	s.dfs(root)
	return s.answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
