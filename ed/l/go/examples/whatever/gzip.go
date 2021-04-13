package main

import (
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

func main() {
	pack("hw.go", "/tmp/")
	unpack("/tmp/hw.go.gz", "/tmp/")
}

func pack(srcFilePath, targetDirPath string) {
	reader, err := os.Open(srcFilePath)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	fileName := filepath.Base(srcFilePath)
	target := filepath.Join(targetDirPath, fileName+".gz")
	writer, err := os.Create(target)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	archiver := gzip.NewWriter(writer)
	archiver.Name = fileName
	defer archiver.Close()

	_, err = io.Copy(archiver, reader)
	if err != nil {
		panic(err)
	}
}

func unpack(srcFilePath, targetDirPath string) {
	reader, err := os.Open(srcFilePath)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	target := filepath.Join(targetDirPath, "res.out.go")
	writer, err := os.Create(target)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)
	if err != nil {
		panic(err)
	}
}
