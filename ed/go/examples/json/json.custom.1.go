package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person string

// (magic)
func (p *Person) UnmarshalJSON(b []byte) error {
	fmt.Printf("<<<MAGIC: p: %+v, b: %+v.\n", string(*p), string(b))

	if strings.HasPrefix(string(*p), "{") {
		value := map[string]string{}
		json.Unmarshal(b, &value)
		*p = Person(value["name"])
	} else {
		*p = Person(b)
	}

	return nil
}

func (p Person) MarshalJSON() ([]byte, error) {
	return []byte(p), nil
}

type PersonStruct struct {
	Person Person `json:"person"`
}

func main() {
	one := `{"person": "john"}`
	two := `{"person": {name:"john"}}`
	result := PersonStruct{}
	json.Unmarshal([]byte(one), &result)
	fmt.Println(result)
	json.Unmarshal([]byte(two), &result)
	fmt.Println(result)
}
