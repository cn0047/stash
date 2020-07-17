package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	f1()
}

func f1() {
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
