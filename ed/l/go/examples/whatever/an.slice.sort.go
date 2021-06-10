package main

import (
	"fmt"
	"sort"
)

func main() {
	sortInt32([]int32{3, -1, 11, 9, 2})
	sortInt32([]int32{3, -7, 0})
}

func sortInt32(arrInt32 []int32) {
	sort.Slice(arrInt32, func(i, j int) bool {
		return arrInt32[i] < arrInt32[j]
	})
	fmt.Println("[] %v \n", arrInt32)
}
