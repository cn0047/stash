package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func issue() {
	f := func(data io.Reader) string {
		d, err := ioutil.ReadAll(data)
		if err != nil {
			panic(err)
		}
		return string(d)
	}

	buf := bytes.NewReader([]byte("msg str"))
	r1 := f(buf)
	fmt.Printf("1a: %v; \n", r1) // 1: msg str;
	r2 := f(buf)
	fmt.Printf("1b: %v; \n", r2) // 2: ;
}

func fix() {
	//tee := io.TeeReader(buf, &t)
}

func main() {
	issue()
	fix()
}
