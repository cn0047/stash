package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	f1()
}

func f1() {
	data := []map[string]interface{}{
		{"id": 1, "k": 1, "v": "one"},
		{"id": 1, "k": 3, "v": "v3"},
	}
	b := make([]byte, 0, len(data))
	for _, el := range data {
		j, err := json.Marshal(el)
		if err != nil {
			panic(err)
		}
		b = append(b, j...)
		b = append(b, []byte("\n")...)
	}
	fmt.Printf("Result: \n%s\n", b)
	err2 := ioutil.WriteFile("/tmp/x.json", b, 0644)
	if err2 != nil {
		panic(err2)
	}
}
