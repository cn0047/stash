package main

import (
	"fmt"
	"bytes"
	"io"
	"encoding/csv"
)

var (

	csvData = `
id,name,email
1,bond,007@mi6.com
2,leiter,leiter@cia.com
3,q,q@mi6.com
`
)

func main() {
	r := bytes.NewReader([]byte(csvData))
	c := NewCSVReader(r)
	for {
		data := c.GetRecord()
		if  data == nil {
			break
		}
		fmt.Printf("%s\n", data)
	}
}

type CSVReader struct {
	reader *csv.Reader
	headers map[string]int
}

func NewCSVReader(r io.Reader) *CSVReader {
	c := CSVReader{}
	c.reader = csv.NewReader(r)
	c.headers = make(map[string]int)
	c.init()
	return &c
}

func (c *CSVReader) init() {
	record, err := c.reader.Read()
	if err == io.EOF {
		panic(fmt.Errorf(": %#v", err))
	}
	if err != nil {
		panic(fmt.Errorf("got unhandled error: %#v", err))
	}

	for i, fieldName := range record {
		c.headers[fieldName] = i
	}
}

func (c CSVReader) GetRecord() map[string]string {
	record, err := c.reader.Read()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		panic(fmt.Errorf("got unhandled error: %s", err))
	}

	data := make(map[string]string)
	for fieldName, i := range c.headers {
		data[fieldName] = record[i]
	}

	return data
}
