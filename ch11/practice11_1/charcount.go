package practice11_1

import (
	"unicode"
	"unicode/utf8"
)

func charCount(s []byte) (counts map[rune]int, utfLen [utf8.UTFMax + 1]int, invalid int) {
	counts = make(map[rune]int)

	for len(s) != 0 {
		r, n := utf8.DecodeRune(s)
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utfLen[n]++

		s = s[n:]
	}
	return
}
