package main

import (
	"fmt"
)

// getHeap returns next heap.
/*
	arr: [10 30 90 100 50]
	tree:
	      10
	     /  \
	   30    90
	  /  \
	100  50
*/
func getHeap() []int {
	h := []int{}
	h = Add(h, 50)
	h = Add(h, 100)
	h = Add(h, 90)
	h = Add(h, 30)
	h = Add(h, 10)

	assert("len", len(h), 5)
	assert("min", h[0], 10)

	return h
}

// heapifyUp performs heapify up.
// Plan:
// get element, in loop get parent,
// if parent has greater value - swap, continue loop.
func heapifyUp(arr []int, i int) []int {
	for {
		p := (i - 1) / 2 // parent
		if p == i || !greaterThan(arr, p, i) {
			break
		}
		arr = swap(arr, p, i)
		i = p
	}

	return arr
}

// heapifyDown performs heapify down.
func heapifyDown(arr []int, idx int, n int) ([]int, bool) {
	i := idx
	for {
		rc := 2*i + 1          // right child
		if rc >= n || rc < 0 { // rc < 0 after int overflow
			break
		}
		lc := rc // left child
		j2 := rc + 1
		if j2 < n && lessThan(arr, j2, rc) {
			lc = j2 // = 2*i + 2  // right child
		}
		if !lessThan(arr, lc, i) {
			break
		}
		arr = swap(arr, i, lc)
		i = lc
	}

	return arr, i > idx
}

func Add(arr []int, v int) []int {
	arr = append(arr, v)
	arr = heapifyUp(arr, len(arr)-1)

	return arr
}

func DelVal(arr []int, v int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == v {
			arr, _ = DelByIndex(arr, i)
		}
	}

	return arr
}

func DelByIndex(arr []int, i int) ([]int, int) {
	n := len(arr) - 1
	if n != i {
		arr = swap(arr, i, n)
		arr, d := heapifyDown(arr, i, n)
		if !d {
			arr = heapifyUp(arr, i)
		}
	}

	v := arr[n]
	arr = arr[0:n]

	return arr, v
}

func DelRoot(arr []int) ([]int, int) {
	n := len(arr) - 1
	arr = swap(arr, 0, n)
	arr, _ = heapifyDown(arr, 0, n)

	v := arr[n]
	arr = arr[0:n]

	return arr, v
}

func lessThan(arr []int, i int, j int) bool {
	return arr[i] < arr[j]
}

func greaterThan(arr []int, i int, j int) bool {
	return arr[i] > arr[j]
}

func swap(arr []int, i int, j int) []int {
	arr[i], arr[j] = arr[j], arr[i]
	return arr
}

func main() {
	heapCheck1()
	heapCheck2()
	heapCheck3()
	heapCheck4()
}

func heapCheck4() {
	h := getHeap()

	h, _ = DelRoot(h)
	assert("len2", len(h), 4)
	assert("min2", h[0], 30)
}

func heapCheck3() {
	h := getHeap()

	h = DelVal(h, 10)
	assert("len2", len(h), 4)
	assert("min2", h[0], 30)

	h = DelVal(h, 90)
	assert("len3", len(h), 3)
	assert("min3", h[0], 30)

	h, _ = DelByIndex(h, 1)
	assert("len4", len(h), 2)
	assert("min4", h[0], 30)
}

func heapCheck2() {
	h := getHeap()

	h, _ = DelRoot(h)
	assert("len2", len(h), 4)
	assert("min2", h[0], 30)
	h, _ = DelRoot(h)
	assert("len3", len(h), 3)
	assert("min3", h[0], 50)
}

func heapCheck1() {
	h := getHeap()

	h, _ = DelRoot(h)
	assert("len2", len(h), 4)
	assert("min2", h[0], 30)

	h, _ = DelRoot(h)
	assert("len3", len(h), 3)
	assert("min3", h[0], 50)
}

func assert(msg string, actual int, expected int) {
	if actual == expected {
		fmt.Printf(".")
		return
	}
	fmt.Printf("[%s] actual: %+v, expected: %v \n", msg, actual, expected)
}
