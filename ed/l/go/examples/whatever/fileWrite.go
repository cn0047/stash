package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	// writeString()
	writeIOReadSeeker()
}

func writeIOReadSeeker() {
	ioReadSeeker("/tmp/debug.txt", strings.NewReader("it works!\n"))
}

func ioReadSeeker(path string, data io.ReadSeeker) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, data)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}

func writeString() {
	fmt.Println("run: tail -f /tmp/debug.log\nTo stop press any key.")

	f, err := os.OpenFile("/tmp/debug.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}

	go writeFile(f)
	_, err = fmt.Scanln()
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}

func writeFile(f *os.File) {
	for {
		time.Sleep(1 * time.Second)
		msg := fmt.Sprintf("dbg log at: %s \n", time.Now().Format(time.Stamp))
		_, err := f.WriteString(msg)
		if err != nil {
			panic(err)
		}
	}
}
