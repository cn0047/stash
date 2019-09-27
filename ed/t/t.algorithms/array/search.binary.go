package main

import (
	"fmt"
)

func main() {
	a := []int{1, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13}
	fmt.Println(binarySearch(a, 6) == 4)
	fmt.Println(binarySearch(a, 9) == 7)
	fmt.Println(binarySearch(a, 7) == 5)
	fmt.Println(binarySearch(a, 5) == 3)
}

func binarySearch(a []int, v int) int {
	start := 0
	end := len(a) - 1
	for start <= end {
		fmt.Print(".")
		mid := (start + end) / 2
		if a[mid] == v {
			return mid
		}
		if a[mid] > v {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}
