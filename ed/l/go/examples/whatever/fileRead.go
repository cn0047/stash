package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	printFileContent("/tmp/debug.txt")
	// printFileContentSafe("/tmp/debug.txt")
}

func printFileContent(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}

func printFileContentSafe(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	buf := make([]byte, 32*1024)

	for {
		n, err := file.Read(buf)
		if n > 0 {
			fmt.Printf("%s", buf[:n])
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Errorf("failed to read %d bytes, error: %v", n, err)
			break
		}
	}
}
