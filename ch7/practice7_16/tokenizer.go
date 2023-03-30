package practice7_16

import (
	"strconv"
)

type token struct {
	number float64
	op     byte
}

func operatorToken(op byte) token {
	return token{
		op: op,
	}
}

func numberToken(num float64) token {
	return token{number: num}
}

func (t token) isNumberToken() bool {
	return t.op == 0
}

type Tokenizer struct {
	buffer []byte
	result []token
}

func NewTokenizer() *Tokenizer {
	var buffer []byte
	var result []token

	tokenizer := Tokenizer{
		buffer: buffer,
		result: result,
	}

	return &tokenizer
}

func (t *Tokenizer) tokenize(input string) []token {
	bytes := []byte(input)

	operators := map[byte]bool{
		'+': true,
		'-': true,
		'*': true,
		'/': true,
	}

	for _, char := range bytes {
		isNumber := '0' <= char && char <= '9'
		if isNumber || char == '.' {
			t.buffer = append(t.buffer, char)
			continue
		}
		if operators[char] {
			t.flushBuffer()
			t.result = append(t.result, operatorToken(char))
		}
	}

	t.flushBuffer()

	return t.result
}

func (t *Tokenizer) flushBuffer() {
	if len(t.buffer) != 0 {
		result, err := strconv.ParseFloat(string(t.buffer), 64)
		if err != nil {
			panic(err)
		}
		t.result = append(t.result, numberToken(result))
		t.buffer = make([]byte, 0)
	}
}
