package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Data struct {
	Name  string `mapstructure:"name"`
	Value string `mapstructure:"value"`
}

func main() {
	a := Data{Name: "test"}
	b := map[string]interface{}{"type": "CLI", "value": "this will be deleted"}

	// Extend map b with data from struct a.
	err := mapstructure.Decode(a, &b)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Final extended map: %#v \n", b)
}
