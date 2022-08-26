// Merge sort (reqursive).
// @see: https://upload.wikimedia.org/wikipedia/commons/c/cc/Merge-sort-example-300px.gif
//
// Worst case = O(n log n),(less than O(n^2)).
// Average case = O(n log n).
// Best case = O(n log n).
// Space required = O(n)
// Provides ability to perform sub-sorts in parallel.
// Predicteble algorithm because only size of array influence performance.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(MergeSort([]int{9, 3, 11, -1, 3, 7}))
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	part1 := MergeSort(arr[mid:])
	part2 := MergeSort(arr[:mid])
	result := merge(part1, part2)

	return result
}

func merge(left, right []int) []int {
	size := len(left) + len(right)
	result := make([]int, size, size)

	leftI := 0
	rightI := 0
	leftLen := len(left) - 1
	rightLen := len(right) - 1

	for i := 0; i < size; i++ {
		if leftI > leftLen && rightI <= rightLen { // Reached end of left slice.
			result[i] = right[rightI]
			rightI++
		} else if leftI <= leftLen && rightI > rightLen { // Reached end of right slice.
			result[i] = left[leftI]
			leftI++
		} else if left[leftI] < right[rightI] { // Use value from left slice.
			result[i] = left[leftI]
			leftI++
		} else { // Use value from right slice.
			result[i] = right[rightI]
			rightI++
		}
	}

	return result
}
