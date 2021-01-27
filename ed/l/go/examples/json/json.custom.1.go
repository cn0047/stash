package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	// checkPersonExample()
	checkDataExample()
}

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

func checkPersonExample() {
	one := `{"person": "john"}`
	two := `{"person": {name:"john"}}`
	result := PersonStruct{}
	json.Unmarshal([]byte(one), &result)
	fmt.Println(result)
	json.Unmarshal([]byte(two), &result)
	fmt.Println(result)
}

type Data struct {
	Name         string `json:"name"`
	ExtendedData string `json:"extended_data"`
	lock         bool
}

func (d *Data) UnmarshalJSON(b []byte) error {
	if d.lock {
		return nil
	}
	d.lock = true
	fmt.Printf("\t %s\n", b)

	err := json.Unmarshal(b, &d)
	d.lock = false
	return err
}

func checkDataExample() {
	j := `{"name":"test", "extended_data":"{\"foo\":\"bar\"}"}`
	d := Data{}
	_ = json.Unmarshal([]byte(j), &d)
	fmt.Printf("%#v\n", d) // main.Data{Name:"", ExtendedData:"", lock:false}
}
