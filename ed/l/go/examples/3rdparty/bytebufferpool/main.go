package main

import (
	"bytes"
	"fmt"
	_ "image/gif"
	"io"

	"github.com/valyala/bytebufferpool"
)

func main() {
	//simpleExample()
	readAfterPutBufferBack()
	//readFromReader()
}

func simpleExample() {
	bb := bytebufferpool.Get()

	_, _ = bb.WriteString("a")
	_, _ = bb.Write([]byte("b"))
	bb.B = append(bb.B, 'c')

	fmt.Printf("buffer: %s\n", bb.B) // buffer: abc

	bytebufferpool.Put(bb)
}

func readAfterPutBufferBack() {
	// Prepare buffer.
	bb := bytebufferpool.Get()
	_, _ = bb.WriteString("a")
	_, _ = bb.WriteString("b")
	buf := bytes.NewReader(bb.B)
	bytebufferpool.Put(bb)

	// Read buffer.
	data, err := io.ReadAll(buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("buffer: %s\n", data)

	buf.Reset(data)
	data2, err := io.ReadAll(buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("buffer: %s\n", data2)
}

func readFromReader() {
	buf := bytes.NewReader([]byte("readFromReader"))

	bb := bytebufferpool.Get()
	_, err := bb.ReadFrom(buf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("buffer: %s\n", bb.B) // buffer: readFromReader

	bytebufferpool.Put(bb)
}

func he(err error) {
	if err != nil {
		panic(err)
	}
}
