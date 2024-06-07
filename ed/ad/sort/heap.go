// Heap sort.
//
// Worst case = O(n log n).
package main

import (
	"fmt"
)

func main() {
	assert(hs([]int{121, 10, 130, 57, 36, 17}))
	assert(hs([]int{9, 3, 11, -1, 3, 7}))
	assert(hs([]int{52, 14, 6, 81, 3, 2, 4, 16, 8, 9, 21, 4, 22, 8, 0}))
	assert(hs([]int{2, 4, 6, 8, 0, -1, 92}))
	assert(hs([]int{2, 4, 6, 8, 3}))
	assert(hs([]int{2, 4, 6, 8, 3}))
	assert(hs([]int{1, -1}))
}

func assert(arr []int) {
	s := "ok"
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			s = "ERR"
		}
	}

	fmt.Printf("%s \t %v \n", s, arr)
}

// hs - heapSort.
func hs(arr []int) []int {
	n := len(arr)

	// Build heap.
	for i := n/2 - 1; i >= 0; i-- { // n/2 - 1 - middle of the tree.
		heapify(arr, i, n)
	}

	// Extract from heap root.
	for i := n - 1; i >= 0; i-- {
		swap(arr, 0, i)
		heapify(arr, 0, i)
	}

	return arr
}

// heapify rearranges elements in slice arr from position i to position n (end),
// it builds "Max heap".
// @uses: recursion.
func heapify(arr []int, i int, n int) {
	root := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && arr[l] > arr[root] {
		root = l
	}
	if r < n && arr[r] > arr[root] {
		root = r
	}

	if root != i {
		swap(arr, i, root)
		heapify(arr, root, n)
	}
}

func swap(arr []int, a int, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
