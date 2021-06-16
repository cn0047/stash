package main

import (
	"encoding/json"
	"fmt"
)

type Req1 struct {
	Val    string `json:"val"`
	Hidden string `json:"-"`
}

type Req2 struct {
	Val    string
	Hidden string
}

type RegsBag struct {
	Regs []Req2
}

func main() {
	// withTags()
	// withoutTags()
	nullStringAfterMarshal()
}

func withTags() {
	p := []byte(`{"val": "yes", "hidden": "no"}`)
	r := Req1{}
	err := json.Unmarshal(p, &r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v \n", r) // main.Req1{Val:"yes", Hidden:""}
}

func withoutTags() {
	p := []byte(`{"val": "yes", "hidden": "no"}`)
	r := Req2{}
	err := json.Unmarshal(p, &r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v \n", r) // main.Req2{Val:"yes", Hidden:"no"}
}

func nullStringAfterMarshal() {
	r := RegsBag{}
	j, err := json.Marshal(r.Regs)
	if err != nil {
		panic(err)
	}

	fmt.Printf("[nullStringAfterMarshal] %s \n", j)
}
