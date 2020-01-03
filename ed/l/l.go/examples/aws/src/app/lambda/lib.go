package main

import (
	"github.com/aws/aws-lambda-go/events"
)

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
