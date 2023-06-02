package main

import (
	"fmt"

	"github.com/grpc/two/lib"

	"github.com/thepkg/strings"
)

type DataBag struct {
	Name string
	Code int
	Meta SomeData
	Data SomeData
}

type SomeData struct {
	Value string
}

type MyServer interface {
	GetByID(id string) (data SomeData, err error)
}

func main() {
	var s any
	s = strings.ToUpperFirst("init")
	s = lib.DataBag{Name: strings.ToUpperFirst("foo")}
	fmt.Printf("%+v \n", s)
}
