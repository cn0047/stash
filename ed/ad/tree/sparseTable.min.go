// @category advanced
package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{4, 6, 8, 7, 3, 2, 9, 5, 1}

	sparseTable := make([][]int, len(arr))
	n := int(math.Ceil(math.Log2(float64(len(arr))))) + 1
	for i := range sparseTable {
		sparseTable[i] = make([]int, n)
	}
	buildSparseTable(arr, sparseTable)

	fmt.Println("Range Minimum Queries (2, 7):", query(2, 7, sparseTable))
	fmt.Println("Range Minimum Queries (0, 2):", query(0, 2, sparseTable))
	fmt.Println("Range Minimum Queries (0, 8):", query(0, 8, sparseTable))
	fmt.Println("Range Minimum Queries (4, 5):", query(4, 5, sparseTable))
	fmt.Println("Range Minimum Queries (7, 8):", query(7, 8, sparseTable))
	fmt.Println("Range Minimum Queries (1, 4):", query(1, 4, sparseTable))
}

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
			sparseTable[r][c] = min(sparseTable[r][c-1], sparseTable[r+(1<<(c-1))][c-1])
			r++
		}
	}
}

func query(left, right int, sparseTable [][]int) int {
	powerOf2 := int(math.Log2(float64(right + 1 - left)))
	i := right + 1 - (1 << powerOf2)

	if sparseTable[left][powerOf2] <= sparseTable[i][powerOf2] {
		i = left
	}

	return sparseTable[i][powerOf2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
