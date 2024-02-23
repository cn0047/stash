package main

import (
	"fmt"
)

type KeyValue struct {
	Key   string
	Value interface{}
}

type Map struct {
	data []KeyValue
}

func (m *Map) Put(key string, value interface{}) {
	for i := range m.data {
		if m.data[i].Key == key {
			m.data[i].Value = value
			return
		}
	}
	m.data = append(m.data, KeyValue{Key: key, Value: value})
}

func (m *Map) Get(key string) (interface{}, bool) {
	for _, kv := range m.data {
		if kv.Key == key {
			return kv.Value, true
		}
	}
	return nil, false
}

func (m *Map) Delete(key string) {
	for i, kv := range m.data {
		if kv.Key == key {
			m.data = append(m.data[:i], m.data[i+1:]...)
			return
		}
	}
}

func main() {
	m := Map{}
	m.Put("f", "foo")
	m.Put("b", "bar")
	m.Delete("b")
	val, ok := m.Get("f")
	fmt.Printf("ok: %v, val: %v \n", ok, val)
}
