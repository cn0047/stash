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

func quickSort(a []int, begin int, end int) []int {
	if begin < end {
		p := partition(a, begin, end)
		a = quickSort(a, begin, p-1)
		a = quickSort(a, p+1, end)
	}

	return a
}

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

func swap(a []int, i int, j int) []int {
	v1 := a[i]
	v2 := a[j]

	a = append(a[:i], append([]int{v2}, a[i+1:]...)...)
	a = append(a[:j], append([]int{v1}, a[j+1:]...)...)

	return a
}
