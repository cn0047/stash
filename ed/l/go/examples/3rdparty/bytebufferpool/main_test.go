package main

import (
	"testing"

	"github.com/valyala/bytebufferpool"
)

func Test_One(t *testing.T) {
	bb := bytebufferpool.Get()
	_, _ = bb.WriteString("test")
	bytebufferpool.Put(bb)
}
