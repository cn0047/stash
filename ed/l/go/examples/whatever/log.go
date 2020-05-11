package main

import (
	"bytes"
	"fmt"
	goLog "log"
)

func main() {
	logger, buf := getDefaultLogger()
	defer func() {
		fmt.Print(buf)
	}()

	f0()
	f1(logger)
	f2(logger) // ⚠️ defer won't be executed
}

func getDefaultLogger() (*goLog.Logger, *bytes.Buffer) {
	var buf bytes.Buffer
	logger := goLog.New(&buf, "", goLog.LstdFlags)

	return logger, &buf
}

func getInfoLogger() (*goLog.Logger, *bytes.Buffer) {
	var buf bytes.Buffer
	logger := goLog.New(&buf, "INFO: ", goLog.Lshortfile)

	return logger, &buf
}

func f0() {
	fmt.Print("[f0] fmt\n") // [f0] fmt
	goLog.Print("[f0] log") // 2020/05/11 09:14:39 [f0] log
}

func f1(log *goLog.Logger) {
	log.Print("[f1]") // 2020/05/11 09:14:39 [f1]
}

func f2(log *goLog.Logger) {
	log.Fatal("[f2]")
}
