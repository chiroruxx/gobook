package main

import (
	"bytes"
	"fmt"
	"os"
)

func comma(s string) string {
	var buf bytes.Buffer

	for i := 0; i < len(s); i++ {
		if i != 0 && (len(s)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}

	return buf.String()
}

func main() {
	fmt.Println(comma(os.Args[1]))
}
