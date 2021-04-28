package main

import (
	"fmt"
	"github.com/valyala/bytebufferpool"
)

func main() {
	f1()
}

func f1() {
	bb := bytebufferpool.Get()

	_, _ = bb.WriteString("line 1\n")
	_, _ = bb.Write([]byte("line 2\n"))
	bb.B = append(bb.B, "line 3\n"...)

	fmt.Printf("buffer:\n%s\n", bb.B)

	bytebufferpool.Put(bb)
}
