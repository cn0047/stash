// GOPATH=$PWD/ed/go/examples/ftp
// go get -u github.com/jlaffaye/ftp

package main

import (
	"github.com/jlaffaye/ftp"
	"io"
	"os"
)

func main() {
	c, err := ftp.Connect("datatransfer.cj.com:21")
	if err != nil {
		panic(err)
	}

	er := c.Login("1994789", "VgYyzUq*")
	if er != nil {
		panic(er)
	}

	c.ChangeDir("/outgoing/productcatalog/215870")

	r, err := c.Retr("Forever_21-Forever21_Google_Feed-shopping.txt.zip")
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
