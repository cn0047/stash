package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	url := "http://realtimelog.herokuapp.com/log"
	filePath := "/Users/k/Downloads/log.html"

	out, err := os.Create(filePath)
	if err != nil {
		println(err)
		return
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		println(err)
		return
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		println(err)
		return
	}
}