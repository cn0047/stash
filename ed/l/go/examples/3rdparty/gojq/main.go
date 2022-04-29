package main

import (
	"encoding/json"
	"fmt"

	"github.com/itchyny/gojq"
)

func main() {
	var err error

	err = jq(`{"list":["a", "b", "c"]}`, ".list[]")
	err = jq(`{"data":{"status":"ok","items":[{"id":1,"name":"foo"},{"id":2,"name":"bar"}]}}`, ".data.items[].name")

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

	fmt.Printf("[jq] Result for: %s | %s \n", input, query)
	i := q.Run(in)
	for {
		v, ok := i.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return fmt.Errorf("got err: %w", err)
		}
		fmt.Printf("%#v\n", v)
	}

	return nil
}
