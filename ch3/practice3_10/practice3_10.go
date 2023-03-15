package practice3_10

import (
	"bytes"
)

func Comma(s string) string {
	var buf bytes.Buffer

	for i := 0; i < len(s); i++ {
		if i != 0 && (len(s)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}

	return buf.String()
}
