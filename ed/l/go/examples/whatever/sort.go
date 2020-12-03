package main

import (
	"fmt"
	"sort"
)

func main() {
	// suint64()
	// sstr()
	sinterface()
}

func suint64() {
	data := []uint64{1, 9, 7, 3}
	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	fmt.Printf("[suint64] data: %+v", data)
}

func sstr() {
	data := []string{"foo", "bar", "a"}
	sort.Strings(data)
	fmt.Printf("[sstr] data: %+v", data)
}

func sinterface() {
	data := make([]interface{}, 0)
	data = append(data, uint64(3))
	data = append(data, uint64(7))
	data = append(data, uint64(1))
	sort.Slice(data, func(i, j int) bool { return data[i].(uint64) < data[j].(uint64) })
	fmt.Printf("[sinterface] data: %+v", data)
}
