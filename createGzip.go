package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	writer.Header.name = "test.txt"
	io.WritterString(writer, "gzip.Writer example\n")
	writer.Close()
}
