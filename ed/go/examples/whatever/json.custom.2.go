package main

import (
	"encoding/json"
	"log"
)

type MyStr struct {
	Scenario string
	Foo string
	Bar string
}

func (m MyStr) MarshalJSON() ([]byte, error) {
	if m.Scenario == "foo" {
		return json.Marshal(struct{ Foo string }{Foo: m.Foo})
	}

	return json.Marshal(struct {Bar string}{Bar: m.Bar})
}

func main() {
	m := MyStr{Scenario: "bar", Foo: "foo", Bar: "bar"}
	j, err := json.Marshal(m)
	if err != nil {
		// error handling
	}
	log.Printf("%s \n", j)
}
