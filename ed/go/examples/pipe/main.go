package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("data.csv")
	check("ERR1: ", err)

	r := bufio.NewReader(f)

	l, _, err := r.ReadLine()
	check("ERR2: ", err)
	write(".header", l)

	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		check("ERR6: ", err)
		writeData(".data.", l)
	}

	er := f.Close()
	check("ERR5: ", er)
}

func check(e string, err error) {
	if err != nil {
		panic(e + err.Error())
	}
}

func write(name string, data []byte) {
	f, _ := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)

	_, err := f.Write(data)
	check("ERR3: ", err)

	er := f.Close()
	check("ERR4: ", er)
}

var (
	n int
	i int
	f *os.File
)

func writeData(name string, data []byte) {
	i++

	if f == nil {
		s := strconv.Itoa(n)
		f, _ = os.OpenFile(name+s, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	}

	_, err := f.Write(append(data, '\n'))
	check("ERR9: ", err)

	if i == 5 {
		err := f.Close()
		check("ERR10: ", err)
		f = nil
		i = 0
		n++
	}
}
