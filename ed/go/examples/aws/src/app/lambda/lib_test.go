package main

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func toAV(t *testing.T, s string) map[string]events.DynamoDBAttributeValue {
	av := events.DynamoDBAttributeValue{}
	err := json.Unmarshal([]byte(s), &av)
	assert.Nil(t, err)
	m := map[string]events.DynamoDBAttributeValue{"av": av}

	return m
}

// @covers getValue
func Test_FromDynamoDBMap(ts *testing.T) {
	ts.Run("Binary", func(t *testing.T) {
		m := FromDynamoDBMap(toAV(t, `{"B": "AAEqQQ=="}`))
		v := m["av"].([]byte)
		assert.Equal(t, "AAEqQQ==", base64.StdEncoding.EncodeToString(v))
	})

	ts.Run("BinarySet", func(t *testing.T) {
		m := FromDynamoDBMap(toAV(t, `{"BS": ["AAEqQQ==", "AAEqQQ=="]}`))
		v := m["av"].([][]byte)[1]
		assert.Equal(t, "AAEqQQ==", base64.StdEncoding.EncodeToString(v))
	})

	ts.Run("Boolean", func(t *testing.T) {
		m := FromDynamoDBMap(toAV(t, `{"BOOL": true}`))
		assert.Equal(t, true, m["av"].(bool))
	})

	ts.Run("List", func(t *testing.T) {
		m := FromDynamoDBMap(toAV(t, `{"L": [{"S": "foo"}, {"S": "bar"}, {"N": "3.14159"}]}`))
		assert.Equal(t, "bar", m["av"].([]interface{})[1])
	})

	ts.Run("List in list", func(t *testing.T) {
		s := `{"L": [{"S": "foo"}, {"S": "bar"}, {"N": "3.14159"}, {"L": [{ "S": "cookies"}]}]}`
		m := FromDynamoDBMap(toAV(t, s))
		v := m["av"].([]interface{})[3].([]interface{})[0]
		assert.Equal(t, "cookies", v)
	})

	ts.Run("Map", func(t *testing.T) {
		m := FromDynamoDBMap(toAV(t, `{"M": {"foo": {"S": "bar"}, "bar": {"S": "foo"}}}`))
		v := m["av"].(map[string]interface{})["foo"]
		assert.Equal(t, "bar", v)
	})

	ts.Run("Map simple", func(t *testing.T) {
		m := FromDynamoDBMap(toAV(t, `{"M": {"foo": {"S": "bar"}}}`))
		v := m["av"].(map[string]interface{})["foo"]
		assert.Equal(t, "bar", v)
	})

	ts.Run("Map in map", func(t *testing.T) {
		s := `{"M": {"bar": {"M": {"name": {"S": "foo"}}}}}`
		m := FromDynamoDBMap(toAV(t, s))
		v := m["av"].(map[string]interface{})["bar"].(map[string]interface{})["name"]
		assert.Equal(t, "foo", v)
	})

	ts.Run("Map in map in map", func(t *testing.T) {
		s := `{"M": {"map": {"M": {"s": {"S": "str"}, "map2": {"M": {"100": {"N": "NOT BLANK"}}}}}}}`
		m := FromDynamoDBMap(toAV(t, s))
		v := m["av"].(map[string]interface{})["map"].(map[string]interface{})["map2"]
		v2 := v.(map[string]interface{})["100"]
		assert.Equal(t, "NOT BLANK", v2)
	})

	ts.Run("Null", func(t *testing.T) {
		m := FromDynamoDBMap(toAV(t, `{"NULL": true}`))
		assert.Equal(t, nil, m["av"])
	})

	ts.Run("Number", func(t *testing.T) {
		m := FromDynamoDBMap(toAV(t, `{ "N": "-123.45"}`))
		assert.Equal(t, "-123.45", m["av"])
	})

	ts.Run("NumberSet", func(t *testing.T) {
		m := FromDynamoDBMap(toAV(t, `{"NS": ["1234", "567.8"]}`))
		assert.Equal(t, "567.8", m["av"].([]string)[1])
	})

	ts.Run("String", func(t *testing.T) {
		m := FromDynamoDBMap(toAV(t, `{"S": "it works"}`))
		assert.Equal(t, "it works", m["av"])
	})

	ts.Run("StringSet", func(t *testing.T) {
		m := FromDynamoDBMap(toAV(t, `{"SS": ["foo", "bar"]}`))
		assert.Equal(t, "bar", m["av"].([]string)[1])
	})

	ts.Run("Complicated test case 1", func(t *testing.T) {
		s := `{"L": [
			{"BOOL":false},
			{"NULL":true},
			{"N":"-256"},
			{"N":"0.999"},
			{"S":""},
			{"SS":["baz"]},
			{"L":[{"S":"ok"}]},
			{"M":{"x":{"S":"y"}}},
			{"NS":["-99"]}
		]}`
		m := FromDynamoDBMap(toAV(t, s))
		v := m["av"].([]interface{})

		assert.Equal(t, false, v[0].(bool))
		assert.Equal(t, nil, v[1])
		assert.Equal(t, "-256", v[2])
		assert.Equal(t, "0.999", v[3])
		assert.Equal(t, "", v[4])
		assert.Equal(t, "baz", v[5].([]string)[0])
		assert.Equal(t, "ok", v[6].([]interface{})[0])
		assert.Equal(t, "y", v[7].(map[string]interface{})["x"])
		assert.Equal(t, "-99", v[8].([]string)[0])
	})
}
