package main

import (
	"github.com/aws/aws-lambda-go/events"
)

func toMap(record map[string]events.DynamoDBAttributeValue) map[string]interface{} {
	resultMap := make(map[string]interface{})

	for key, rec := range record {
		switch rec.DataType() {
		case events.DataTypeBinary:
		case events.DataTypeBinarySet:
		case events.DataTypeBoolean:
			resultMap[key] = rec.Boolean()
		case events.DataTypeList:
		case events.DataTypeMap:
		case events.DataTypeNull:
		case events.DataTypeNumber:
			resultMap[key] = rec.Number()
		case events.DataTypeNumberSet:
		case events.DataTypeString:
			resultMap[key] = rec.String()
		case events.DataTypeStringSet:
		}
	}

	return resultMap
}

func ToMapFromDynamoDB(record map[string]events.DynamoDBAttributeValue) map[string]interface{} {
	resultMap := make(map[string]interface{})

	for key, rec := range record {
		switch rec.DataType() {
		case events.DataTypeBinary:
		case events.DataTypeBinarySet:
		case events.DataTypeBoolean:
			resultMap[key] = rec.Boolean()
		case events.DataTypeList:
		case events.DataTypeMap:
		case events.DataTypeNull:
		case events.DataTypeNumber:
			resultMap[key] = rec.Number()
		case events.DataTypeNumberSet:
		case events.DataTypeString:
			resultMap[key] = rec.String()
		case events.DataTypeStringSet:
		}
	}

	return resultMap
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
