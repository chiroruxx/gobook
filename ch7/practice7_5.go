package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type internalLimitReader struct {
	limit  int64
	reader *io.Reader
}

func (reader *internalLimitReader) Read(p []byte) (int, error) {
	originalReader := *reader.reader
	readSize, err := originalReader.Read(p)
	if readSize >= int(reader.limit) {
		p = p[:reader.limit]
		reader.limit = 0
		return readSize, io.EOF
	}

	reader.limit -= int64(readSize)

	return readSize, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	limitReader := new(internalLimitReader)
	limitReader.reader = &r
	limitReader.limit = n

	return limitReader
}

func main() {
	p := make([]byte, 10)

	reader, err := os.Open("bytecounter.go")
	defer reader.Close()
	if err != nil {
		log.Fatalf("limitreader: %v\n", err)
	}

	limitReader := LimitReader(reader, 10)

	fmt.Println(limitReader.Read(p))
	fmt.Println(p)
}
