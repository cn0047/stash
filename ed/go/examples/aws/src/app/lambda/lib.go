package main

import (
	"github.com/aws/aws-lambda-go/events"
)

func FromDynamoDBMap(record map[string]events.DynamoDBAttributeValue) map[string]interface{} {
	resultMap := make(map[string]interface{})

	for key, rec := range record {
		resultMap[key] = getValue(rec)
	}

	return resultMap
}

func getValue(record events.DynamoDBAttributeValue) interface{} {
	var val interface{}

	switch record.DataType() {
	case events.DataTypeBinary:
		val = record.Binary()
	case events.DataTypeBinarySet:
		val = record.BinarySet()
	case events.DataTypeBoolean:
		val = record.Boolean()
	case events.DataTypeList:
		s := make([]interface{}, 0)
		for _, el := range record.List() {
			s = append(s, getValue(el))
		}
		val = s
	case events.DataTypeMap:
		// IMPORTANT: For DynamoDB only string can be a key in map.
		m := make(map[string]interface{})
		for k, el := range record.Map() {
			m[k] = getValue(el)
		}
		val = m
	case events.DataTypeNull:
		val = nil
	case events.DataTypeNumber:
		val = record.Number()
	case events.DataTypeNumberSet:
		val = record.NumberSet()
	case events.DataTypeString:
		val = record.String()
	case events.DataTypeStringSet:
		val = record.StringSet()
	}

	return val
}

func GetStringByKey(key string, record map[string]events.DynamoDBAttributeValue) string {
	v, ok := record[key]
	if !ok {
		return ""
	}

	return v.String()
}

func GetBooleanByKey(key string, record map[string]events.DynamoDBAttributeValue) bool {
	v, ok := record[key]
	if !ok {
		return false
	}

	return v.Boolean()
}

func GetIntegerByKey(key string, record map[string]events.DynamoDBAttributeValue) int64 {
	v, ok := record[key]
	if !ok {
		return 0
	}

	val, err := v.Integer()
	if err != nil {
		return 0
	}

	return val
}
