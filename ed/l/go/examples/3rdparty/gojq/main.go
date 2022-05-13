package main

import (
	"encoding/json"
	"fmt"

	"github.com/itchyny/gojq"
)

const (
	j = `{
		"active": true,
		"extra": null,
		"list": ["a", "b", "c"],
		"data": {
			"status": "ok",
			"items": [{"id":1,"name":"foo","f":1.5}, {"id":2,"name":"bar","f":0.1357}]
		}
	}`
)

func main() {
	var err error

	//err = jq(j, ".active")
	//check(err)
	//err = jq(j, ".extra")
	//check(err)
	//err = jq(j, ".list[]")
	//check(err)
	err = jq(j, ".data")
	check(err)
	//err = jq(j, ".data.items")
	//check(err)
	//err = jq(j, ".data.items[0].id")
	//check(err)
	//err = jq(j, ".data.items[1].f")
	//check(err)
	//err = jq(j, `.data.items[] | {"i":.id, "n":".name"}`)

	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Printf("err: %v \n", err)
	}
}

func toMap(input string) (map[string]interface{}, error) {
	var r map[string]interface{}
	err := json.Unmarshal([]byte(input), &r)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal input, err: %w", err)
	}

	return r, nil
}

func jq(input string, query string) error {
	in, err := toMap(input)
	if err != nil {
		return fmt.Errorf("failed to parse input, err: %w", err)
	}

	q, err := gojq.Parse(query)
	if err != nil {
		return fmt.Errorf("failed to parse query, err: %w", err)
	}

	fmt.Printf("\n[jq] Result for: %s \n", query)
	i := q.Run(in)
	for {
		v, ok := i.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return fmt.Errorf("got err: %w", err)
		}
		err := fromInterface(v)
		if err != nil {
			return fmt.Errorf("failed to perform fromInterface, err: %w", err)
		}
	}

	return nil
}

func fromInterface(v interface{}) error {
	switch t := v.(type) {
	case float64:
		fmt.Printf("num: %#v\n", t)
	case string:
		fmt.Printf("str: %#v\n", t)
	case bool:
		fmt.Printf("bool: %#v\n", t)
	case []interface{}:
		fmt.Printf("array: \n")
		for _, val := range t {
			err := fromInterface(val)
			if err != nil {
				return err
			}
		}
	case map[string]interface{}:
		fmt.Printf("object: \n")
		for k, val := range t {
			fmt.Printf("key: %v; val: ", k)
			err := fromInterface(val)
			if err != nil {
				return err
			}
		}
	case nil:
		fmt.Printf("nil: %#v\n", t)
	default:
		return fmt.Errorf("failed to convert val: %#v with type: %T", v, t)
	}

	return nil
}
