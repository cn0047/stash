// GOPATH=$PWD/ed/go/examples/ftp
// go get -u github.com/jlaffaye/ftp

package main

import (
	"io"
	"os"

	"github.com/jlaffaye/ftp"
)

func main() {
	c, err := ftp.Connect("{host}.com:21")
	if err != nil {
		panic(err)
	}

	er := c.Login("usr", "pwd")
	if er != nil {
		panic(er)
	}

	c.ChangeDir("/tmp")

	r, err := c.Retr("file.txt")
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("/tmp/x.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	_, er2 := io.Copy(file, r)
	if er2 != nil {
		panic(er2)
	}
	file.Close()
}
