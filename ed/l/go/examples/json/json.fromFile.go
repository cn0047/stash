package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f1("/tmp/x2.jsonx")
}

type Record struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func f1(filename string) {
	jsonFile, err := os.Open(filename)
	defer jsonFile.Close()
	if err != nil {
		fmt.Printf("failed to open json file: %s, error: %v", filename, err)
		return
	}

	// data, err := ioutil.ReadAll(jsonFile)
	data, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("failed to read json file, error: %v", err)
		return
	}

	r := Record{}
	if err := json.Unmarshal(data, &r); err != nil {
		fmt.Printf("failed to unmarshal json file, error: %v", err)
		return
	}

	fmt.Printf("Result: \n%+v\n", r)
}
