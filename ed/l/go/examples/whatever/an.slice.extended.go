package main

import (
	"fmt"
)

type Item struct {
	Value int
}

func Set(items []Item, index int, value int) []Item {
	if len(items) < index+1 {
		delta := (index + 1) - len(items)
		items = append(items, make([]Item, delta)...)
	}

	if items[index] == (Item{}) {
		items[index] = Item{}
	}

	items[index].Value = value

	return items
}

func main() {
	setValueToSliceItemByIndexInSlice()
}

func setValueToSliceItemByIndexInSlice() {
	items := make([]Item, 0)
	items = Set(items, 7, 17)
	items = Set(items, 0, 10)
	items = Set(items, 3, 13)
	items = Set(items, 9, 19)

	fmt.Printf("🎾 %+v \n", items)
}
