package main

import (
	"fmt"
)

func main() {
	fmt.Println(hs([]int{121, 10, 130, 57, 36, 17}))
	fmt.Println(hs([]int{9, 3, 11, -1, 3, 7}))
	fmt.Println(hs([]int{52, 14, 6, 81, 3, 2, 4, 16, 8, 9, 21, 4, 22, 8, 0}))
	fmt.Println(hs([]int{2, 4, 6, 8, 0, -1, 92}))
	fmt.Println(hs([]int{2, 4, 6, 8, 3}))
}

func hs(arr []int) []int {
	n := len(arr)
	return heapSort(arr, n)
}

func heapSort(arr []int, n int) []int {
	// build heap
	for i := n/2 - 1; i >= 0; i-- {
		arr = heapify(arr, n, i)
	}
	// extract from heap root
	for i := n - 1; i >= 0; i-- {
		arr = swap(arr, 0, i)
		arr = heapify(arr, i, 0)
	}
	return arr
}

func heapify(arr []int, n int, i int) []int {
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
		arr = swap(arr, i, root)
		arr = heapify(arr, n, root)
	}
	return arr
}

func swap(a []int, i int, j int) []int {
	v1 := a[i]
	v2 := a[j]

	a = append(a[:i], append([]int{v2}, a[i+1:]...)...)
	a = append(a[:j], append([]int{v1}, a[j+1:]...)...)

	return a
}
