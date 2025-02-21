package main

import (
	"fmt"
	"runtime"
	"weak"
)

type Blob []byte

func (b Blob) String() string {
	return fmt.Sprintf("Blob size: %d Kb", len(b)/1024)
}

func newBlob(size int) *Blob {
	b := make([]byte, size*1024)
	for i := range size {
		b[i] = byte(i) % 255
	}
	return (*Blob)(&b)
}

func main() {
	b := newBlob(1000)
	fmt.Println(b) // Blob size: 1000 Kb

	wb := weak.Make(newBlob(1000))
	fmt.Println(wb.Value()) // Blob size: 1000 Kb

	runtime.GC()

	fmt.Println(b)          // Blob size: 1000 Kb
	fmt.Println(wb.Value()) // nil
}
