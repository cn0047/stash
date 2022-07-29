// Quick sort (reqursive).
// @see: https://upload.wikimedia.org/wikipedia/commons/6/6a/Sorting_quicksort_anim.gif
//
// Worst case = O(n^2).
// Average case = O(n log n).
// Best case = O(n log n).
// Space required = O(n).
package main

import (
	"fmt"
)

func main() {
	//a := []int{2, 4, 6, 8, 0}
	//a := []int{5, 3, 1, 2, 4}
	//a := []int{2, 4, 6, 8, 3}
	//a := []int{2, 4, 6, 8, 9}
	//a := []int{2, 4, 6, 8, 0, -1, 92}
	a := []int{52, 14, 6, 81, 3, 2, 4, 16, 8, 9, 21, 4, 22, 8, 0}
	n := len(a) - 1
	fmt.Println(quickSort(a, 0, n))
}

func quickSort(arr []int, begin int, end int) []int {
	if begin < end {
		p := partition(arr, begin, end)
		arr = quickSort(arr, begin, p-1)
		arr = quickSort(arr, p+1, end)
	}

	return arr
}

// partition return pivot element position.
func partition(a []int, begin int, end int) int {
	pivot := a[end]
	i := begin - 1
	for j := begin; j < end; j++ {
		if a[j] <= pivot {
			i++
			a = swap(a, i, j)
		}
	}
	swap(a, i+1, end)

	return i + 1
}

func swap(arr []int, a int, b int) []int {
	arr[a], arr[b] = arr[b], arr[a]
	return arr
}
