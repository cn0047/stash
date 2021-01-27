package main

import (
	"encoding/json"
	"fmt"
)

type Entry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	j1 := `{"id":"7"}`
	j2 := `{"name":"seven"}`
	e := Entry{}
	_ = json.Unmarshal([]byte(j1), &e)
	fmt.Printf("1) %+v \n", e)
	_ = json.Unmarshal([]byte(j2), &e)
	fmt.Printf("2) %+v \n", e)
}

/*
1) {ID:7 Name:}
2) {ID:7 Name:seven}
*/
