// @see: https://leetcode.com/problems/range-sum-query-mutable
// @category advanced
package main

import (
	"fmt"
)

// Tree represents segment tree.
type Tree struct {
	n    int
	nums []int // stores actual values.
	tree []int // stores information about segments (ranges).
}

func Constructor(nums []int) Tree {
	n := len(nums)
	internalTree := make([]int, n+1)
	t := Tree{n: n, nums: nums, tree: internalTree}

	for i, val := range nums {
		t.updateTree(i+1, val)
	}

	return t
}

func (t *Tree) updateTree(i int, val int) {
	for i <= t.n {
		t.tree[i] += val
		i += i & -i // index to next node
	}
}

func (t *Tree) getSum(i int) int {
	sum := 0
	for i > 0 {
		sum += t.tree[i]
		i -= i & -i // index to previous node
	}
	return sum
}

func (t *Tree) Update(i int, val int) {
	diff := val - t.nums[i]
	t.nums[i] = val
	t.updateTree(i+1, diff)
}

func (t *Tree) SumRange(left int, right int) int {
	return t.getSum(right+1) - t.getSum(left)
}

func main() {
	nums := []int{}
	t := Tree{}

	nums = []int{1, 3, 5}
	t = Constructor(nums)
	fmt.Println(t.SumRange(0, 2)) // 9
	t.Update(1, 2)
	fmt.Println(t.SumRange(0, 2)) // 8

	nums = []int{5, 8, 7, 2, 10, 2, 1}
	t = Constructor(nums)
	fmt.Println(t.SumRange(1, 5))
}
