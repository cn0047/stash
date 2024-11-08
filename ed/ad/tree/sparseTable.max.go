// @see: https://leetcode.com/problems/sliding-window-maximum
// @category advanced
package main

import (
	"fmt"
	"math"
)

func buildSparseTable(arr []int, sparseTable [][]int) {
	rows := len(arr)
	cols := len(sparseTable[0])

	for r := 0; r < rows; r++ {
		sparseTable[r][0] = arr[r]
	}

	for c := 1; c < cols; c++ {
		shift := 1 << c
		r := 0
		for r+shift <= rows {
			sparseTable[r][c] = max(sparseTable[r][c-1], sparseTable[r+(1<<(c-1))][c-1])
			r++
		}
	}
}

func query(left, right int, sparseTable [][]int) int {
	powerOf2 := int(math.Log2(float64(right + 1 - left)))
	i := right + 1 - (1 << powerOf2)

	if sparseTable[left][powerOf2] >= sparseTable[i][powerOf2] {
		i = left
	}

	return sparseTable[i][powerOf2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSlidingWindow(nums []int, k int) []int {
	sparseTable := make([][]int, len(nums))
	n := int(math.Ceil(math.Log2(float64(len(nums))))) + 1
	for i := range sparseTable {
		sparseTable[i] = make([]int, n)
	}
	buildSparseTable(nums, sparseTable)

	res := make([]int, 0, len(nums)-k)
	for i := 0; i <= len(nums)-k; i++ {
		res = append(res, query(i, i+k-1, sparseTable))
	}

	return res
}

func main() {
	arr, k := []int{1, 3, -1, -3, 5, 3, 6, 7}, 3
	r := maxSlidingWindow(arr, k)
	fmt.Printf("===\n%v", r)
}
