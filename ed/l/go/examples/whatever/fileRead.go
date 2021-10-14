package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// printFileContent("/tmp/debug.txt")
	// printFileContentSafe("/tmp/debug.txt")
	printFileContentSafeLineByLine("/tmp/debug.txt")
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
			panic(fmt.Errorf("failed to read %d bytes, error: %v", n, err))
		}
	}
}

func printFileContentSafeLineByLine(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(fmt.Errorf("failed to read line, error: %v", err))
		}
		fmt.Printf("%s ⮐ \n", line)
	}
}
