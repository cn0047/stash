package main

import (
	"github.com/thepkg/strings"

	"mdl/foo"
	"mdl/foo/bar"
)

func main() {
	println(strings.ToUpperFirst("ok!"))
	println(foo.GetStr("yes, all ok!"))
	println(bar.GetStr("and it's true!"))
}
