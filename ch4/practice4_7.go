package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseUtf8(s []byte) []byte {
	var runes []rune
	var sizes []int
	for index, size := 0, 0; index < len(s); index += size {
		var r rune
		r, size = utf8.DecodeRune(s[index:])

		runes = append(runes, r)
		sizes = append(sizes, size)
	}

	var result []byte
	for index := len(runes) - 1; index >= 0; index-- {
		bytes := make([]byte, sizes[index])
		utf8.EncodeRune(bytes, runes[index])
		result = append(result, bytes...)
	}

	return result
}

func main() {
	s := "こんにちは"
	fmt.Printf("%s\n", s)
	fmt.Printf("%s", reverseUtf8([]byte(s)))
}
