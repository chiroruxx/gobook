package lexer

type Token interface {
	String() string
}

type NumberToken struct {
	numbers []byte
}

func NewNumberToken(numbers []byte) *NumberToken {
	return &NumberToken{numbers: numbers}
}

func (t *NumberToken) String() string {
	return string(t.numbers)
}

type StringToken struct {
	chars []byte
}

func NewStringToken(chars []byte) *StringToken {
	return &StringToken{chars: chars}
}

func (t *StringToken) String() string {
	return string(t.chars)
}

type QuoteToken struct{}

func (q *QuoteToken) String() string {
	return `"`
}

type ListStartToken struct{}

func (l *ListStartToken) String() string {
	return "("
}

type ListEndToken struct{}

func (l *ListEndToken) String() string {
	return ")"
}

type ListSeparatorToken struct{}

func (l *ListSeparatorToken) String() string {
	return " "
}

type EscapeToken struct{}

func (e *EscapeToken) String() string {
	return "\\"
}

type EscapedStringToken struct {
	chr byte
}

func (e *EscapedStringToken) String() string {
	return string(e.chr)
}
