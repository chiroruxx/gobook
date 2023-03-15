package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	count := bytes.Count(p, []byte("\n")) + 1
	*c += LineCounter(count)
	return count, nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	count := 0
	for next := 0; next < len(p); count++ {
		next, _, err := bufio.ScanWords(p, true)
		if err != nil {
			return 0, err
		}
		p = p[next:]
	}
	*c += WordCounter(count)
	return count, nil
}

func main() {
	var l LineCounter
	l.Write([]byte("hello\nhello!!"))
	fmt.Println(l)

	l = 0
	var name = "Dolly"
	fmt.Fprintf(&l, "hello, %s", name)
	fmt.Println(l)

	var w WordCounter
	w.Write([]byte("hello hello!!"))
	fmt.Println(w)

	w = 0
	fmt.Fprintf(&w, "hello")
	fmt.Println(w)

	w = 0
	fmt.Fprintf(&w, "hello, %s", name)
	fmt.Println(w)

	w = 0
	w.Write([]byte("hello hello ag!!"))
	fmt.Println(w)
}
