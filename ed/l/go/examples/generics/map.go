package main

import (
	"fmt"
	"strconv"
)

func main() {
	r1, err := convertKeysToInt64(map[string]string{"1": "10", "2": "20"})
	fmt.Printf("[case1] %+v \t %#v \n", err, r1)

	r2, err := convertKeysToInt64(map[string]int64{"1": 11, "2": 22})
	fmt.Printf("[case2] %+v \t %#v \n", err, r2)
}

func convertKeysToInt64[K string, V int64 | string](input map[K]V) (output map[int64]V, err error) {
	output = make(map[int64]V, len(input))

	for k, v := range input {
		key, err := strconv.ParseInt(string(k), 10, 64)
		if err != nil {
			return output, fmt.Errorf("failed to convert key: %v", k)
		}
		output[key] = v
	}

	return output, nil
}
