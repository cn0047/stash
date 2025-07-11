package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Bag struct {
	Nested    Nested         `json:"bar,inline"`
	Extra     string         `json:"extra,omitempty"`
	Count     int            `json:"count,omitzero"`
	Fraction  float32        `json:"fraction,string"`
	Hidden    string         `json:"-"`
	Data      map[string]any `json:",unknown"`
	UpdatedAt time.Time      `json:"updated_at,format:DateOnly"`
	Case1     string         `json:"case1,case:ignore"`
	Case2     string         `json:"case2,case:strict"`
}

type Nested struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func main() {
	marshal()
	unmarshal()
}

func marshal() {
	v := Bag{
		Nested:    Nested{Code: 2, Text: "bar"},
		Extra:     "",
		Count:     0,
		Fraction:  9.99,
		Hidden:    "hidden value",
		UpdatedAt: time.Now(),
	}
	j, _ := json.Marshal(v)
	fmt.Printf("marshal result: %s \n", j) // {"bar":{"code":2,"text":"bar"},"fraction":"9.99"}
}

func unmarshal() {
	j := `{
		"extra": "json->go",
		"CASE1": "IMPORTANT",
		"CASE2": "ImpOrtAnT",
		"payload": {
			"arr": [{"foo": "bar"}, {"code": 200}]
		}
	}`
	v := &Bag{}
	json.Unmarshal([]byte(j), v)
	fmt.Printf("unmarshal result: %+v \n", v)
}
