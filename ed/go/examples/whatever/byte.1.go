package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

func main() {
	one()
	two()
	three()
	forth()
}

func one() {
	origin := `{"msg": "test"}`

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(origin)
	fmt.Printf("\n %+v", buf.Bytes())

	dec := gob.NewDecoder(&buf)
	var res string
	dec.Decode(&res)
	fmt.Printf("\n %+v", res)
}

func two() {
	origin := map[string]interface{}{"code": 200, "msg": "OK"}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(origin)
	fmt.Printf("\n %+v", buf.Bytes())

	dec := gob.NewDecoder(&buf)
	res := make(map[string]interface{})
	dec.Decode(&res)
	fmt.Printf("\n %+v", res)
}

func three() {
	origin := []interface{}{204, "blank"}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(origin)
	fmt.Printf("\n %+v", buf.Bytes())

	dec := gob.NewDecoder(&buf)
	res := make([]interface{}, 0)
	dec.Decode(&res)
	fmt.Printf("\n %+v", res)
}

func forth() {
	origin := map[string]interface{}{"key": "n", "value": 204}

	enc, _ := json.Marshal(origin)
	fmt.Printf("\n %+v", enc)

	res := make(map[string]interface{})
	json.Unmarshal(enc, &res)
	fmt.Printf("\n %+v", res)
}
