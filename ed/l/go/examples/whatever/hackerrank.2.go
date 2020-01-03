package main

import (
	"bufio"
	"fmt"
	"os"
	// "regexp"
	"strings"
)

func f(s string) string {
	return s
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)
	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	ss := ""
	for {
		s, e := readLine(reader)
		if e != nil {
			break
		}
		ss += "\n" + s
	}

	res := f(ss)
	fmt.Fprintf(writer, "%v", res)

	writer.Flush()
}

func readLine(reader *bufio.Reader) (string, error) {
	str, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}

	return strings.TrimRight(string(str), "\r\n"), nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
