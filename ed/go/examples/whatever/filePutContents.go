package main

import (
	"os"
)

func main() {
	f, _ := os.OpenFile("/tmp/debug.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777); f.WriteString("dbg" + "\n")
}
