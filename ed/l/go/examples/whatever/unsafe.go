package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// sizeOfInt()
	// sizeOfStruct()
	// pointer()
}

type Foo struct {
	Value string
}

func sizeOfInt() {
	fmt.Printf(
		"[sizeOfInt] %v %v %v %v %v \n",
		unsafe.Sizeof(uint(0)),   // 8
		unsafe.Sizeof(uint8(0)),  // 1
		unsafe.Sizeof(uint16(0)), // 2
		unsafe.Sizeof(uint32(0)), // 4
		unsafe.Sizeof(uint64(0)), // 8
	)
}

func sizeOfStruct() {
	f := Foo{Value: "foo"}
	fmt.Printf("[sizeOfStruct] Size of struct: %v \n", unsafe.Sizeof(f)) // Size of struct: 16
}

func pointer() {
	var f1 float64 = 10
	var i1 uint64 = 0
	i1 = *(*uint64)(unsafe.Pointer(&f1))
	fmt.Printf("[] %d \n", i1)

	var i2 uint64 = 10
	var f2 float64 = 0
	f2 = *(*float64)(unsafe.Pointer(&i2))
	fmt.Printf("[] %f \n", f2)
}
