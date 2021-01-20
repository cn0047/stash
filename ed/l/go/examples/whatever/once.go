package main

import (
	"sync"
)

var (
	once sync.Once
	v    int
)

func main() {
	printVValue()
	printVValue()
	printVValue()
}

func incVValue() {
	v++
}

func printVValue() {
	once.Do(incVValue)
	println(v)
}

/*
1
1
1
*/
