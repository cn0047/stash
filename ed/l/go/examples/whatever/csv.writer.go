package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	simpleWrite()
}

func simpleWrite() {
	file, err := ioutil.TempFile("/tmp", "csv")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create temp file, err: %w", err))
	}
	//defer os.Remove(file.Name())

	data := [][]string{
		{"id", "msg"},
		{"1", "test"},
		{"2", "cli"},
		{"3", "ok"},
	}
	err = writeToCSVFile(file.Name(), data)
	fmt.Printf("[simpleWrite] file: %+v, err: %+v \n", file.Name(), err)
}

func writeToCSVFile(filePath string, records [][]string) error {
	csvFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed creating file: %w", err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	for _, r := range records {
		err = writer.Write(r)
		if err != nil {
			return fmt.Errorf("failed to write record, err: %w", err)
		}
	}
	writer.Flush()

	return nil
}
