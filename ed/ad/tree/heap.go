package main

import (
	"fmt"
)

func main() {
	//heapCheck1()
	//heapCheck2()
	//heapCheck3()
	heapCheck4()
}

func heapCheck4() {
	/*
		arr: [10 30 90 100 50]
		tree:
		      10
		     /  \
		   30    90
		  /  \
		100  50
	*/
	h := getHeap()

	h, _ = Del(h)
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

	h, _ = Del(h)
	assert("len2", len(h), 4)
	assert("min2", h[0], 30)
	h, _ = Del(h)
	assert("len3", len(h), 3)
	assert("min3", h[0], 50)
}

func heapCheck1() {
	h := getHeap()

	h, _ = Del(h)
	assert("len2", len(h), 4)
	assert("min2", h[0], 30)

	h, _ = Del(h)
	assert("len3", len(h), 3)
	assert("min3", h[0], 50)
}

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

func assert(msg string, actual int, expected int) {
	if actual == expected {
		return
	}
	fmt.Printf("[%s] actual: %+v, expected: %v \n", msg, actual, expected)
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
		arr, d := down(arr, i, n)
		if !d {
			arr = up(arr, i)
		}
	}

	x := arr[n]
	arr = arr[0:n]

	return arr, x
}

func Del(arr []int) ([]int, int) {
	n := len(arr) - 1
	arr = swap(arr, 0, n)
	arr, _ = down(arr, 0, n)

	x := arr[n]
	arr = arr[0:n]

	return arr, x
}

func down(arr []int, idx int, n int) ([]int, bool) {
	i := idx
	for {
		j1 := 2*i + 1          // right child
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		j2 := j1 + 1
		if j2 < n && less(arr, j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !less(arr, j, i) {
			break
		}
		arr = swap(arr, i, j)
		i = j
	}

	return arr, i > idx
}

func Add(arr []int, v int) []int {
	arr = append(arr, v)
	arr = up(arr, len(arr)-1)

	return arr
}

func up(arr []int, i int) []int {
	for {
		p := (i - 1) / 2 // parent
		if p == i || !less(arr, i, p) {
			break
		}
		arr = swap(arr, p, i)
		i = p
	}

	return arr
}

func less(arr []int, i int, j int) bool {
	return arr[i] < arr[j]
}

func swap(arr []int, i int, j int) []int {
	arr[i], arr[j] = arr[j], arr[i]

	return arr
}
