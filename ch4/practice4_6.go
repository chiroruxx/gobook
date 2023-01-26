package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

const asciiSpace = ' '

func removeSpace(s []byte) []byte {
	i := 0
	var previous rune

	for index, next := 0, 0; index < len(s); index += next {
		r, runeWidth := utf8.DecodeRune(s[index:])
		next = runeWidth

		if unicode.IsSpace(r) {
			r, runeWidth = asciiSpace, 1
			s[index] = byte(r)
		}

		if index == 0 || r != asciiSpace || r != previous {
			for j := 0; j < runeWidth; j++ {
				s[i] = s[index+j]
				i++
			}
			previous = r
		}
	}

	return s[:i]
}

func main() {
	s := "こんにちは　 　！！"
	fmt.Printf("%s\n", s)
	fmt.Printf("%s", removeSpace([]byte(s)))
}
