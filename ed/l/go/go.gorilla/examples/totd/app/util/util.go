package util

import (
	"fmt"
	"strconv"
)

// ConvertToMapWithInt64Keys returns map[int64]V out of interface{}.
func ConvertToMapWithInt64Keys[K string, V int64 | string](input interface{}) (output map[int64]V, err error) {
	output = make(map[int64]V)

	if input == nil {
		return output, nil
	}
	rawMap, ok := input.(map[string]interface{})
	if !ok {
		return output, fmt.Errorf("failed to convert input into map: %v", input)
	}

	for k, v := range rawMap {
		key, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			return output, fmt.Errorf("failed to convert key: %v", k)
		}

		output[key] = v.(V)
	}

	return output, nil
}

// ConvertKeysToInt64 returns map with converted keys to int64.
func ConvertKeysToInt64[K string, V int64 | string](input map[K]V) (output map[int64]V, err error) {
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
