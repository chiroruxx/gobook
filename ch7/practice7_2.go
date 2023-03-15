package main

import (
	"bytes"
	"fmt"
	"io"
)

type internalWriter struct {
	writer  io.Writer
	counter int64
}

func (writer *internalWriter) Write(p []byte) (int, error) {
	writer.counter += int64(len(p))
	return writer.writer.Write(p)
}

func CountingWriter(writer io.Writer) (io.Writer, *int64) {
	internalWriter := internalWriter{writer: writer}

	return &internalWriter, &internalWriter.counter
}

func main() {
	buffer := bytes.Buffer{}
	writer, count := CountingWriter(&buffer)
	fmt.Println(*count)
	writer.Write([]byte{'a'})
	fmt.Println(*count)
	writer.Write([]byte{'a', 'b'})
	fmt.Println(*count)
}
