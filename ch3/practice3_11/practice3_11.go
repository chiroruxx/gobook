package practice3_11

import (
	"bytes"
	"strings"
)

func Comma(s string) string {
	split := strings.Split(s, ".")

	var buf bytes.Buffer
	commaInt(split[0], &buf)

	if len(split) == 2 {
		buf.WriteByte('.')
		commaDecimal(split[1], &buf)
	}

	return buf.String()
}

func commaInt(s string, buf *bytes.Buffer) {
	for i := 0; i < len(s); i++ {
		if i != 0 && (len(s)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
}

func commaDecimal(s string, buf *bytes.Buffer) {
	for i := 0; i < len(s); i++ {
		if i != 0 && i%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
}
