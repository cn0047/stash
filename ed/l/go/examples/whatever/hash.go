package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {
	// simple()

	r, err := GetFileMD5Hash("/tmp/x.json")
	fmt.Printf("err: %v, r: %v \n", err, r)
}

func simple() {
	s := "my string"

	h := fmt.Sprintf("%x", md5.Sum([]byte(s)))
	fmt.Printf("md5 = %s \n", h)

	sh1 := sha1.New()
	sh1.Write([]byte(s))
	fmt.Printf("sha1 = %x \n", sh1.Sum(nil))

	sh256 := sha256.New()
	sh256.Write([]byte(s))
	fmt.Printf("sha256 = %x \n", sh256.Sum(nil))
}

func GetFileMD5Hash(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to open file, err: %v", err)
	}
	defer f.Close()

	h := md5.New()

	_, err = io.Copy(h, f)
	if err != nil {
		return "", fmt.Errorf("failed to copy file, err: %v", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
