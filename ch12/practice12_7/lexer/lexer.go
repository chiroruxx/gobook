package lexer

import (
	"errors"
	"io"
)

func Lex(r io.Reader) ([]Token, error) {
	l := newLexer(r)
	return l.Lex()
}

type byteType uint

const (
	typeInitial byteType = iota
	typeNumber
	typeQuote
	typeString
	typeListStart
	typeListEnd
	typeListSeparator
	typeEscape
	typeEscapedString
)

func (t byteType) canNest() bool {
	return t == typeListStart || t == typeListEnd
}

type lexer struct {
	r       io.Reader
	typ     byteType
	current []byte
}

func newLexer(r io.Reader) *lexer {
	return &lexer{
		r: r,
	}
}

func (l *lexer) Lex() ([]Token, error) {
	l.typ = typeInitial

	var tokens []Token
	for {
		b, _, err := l.readByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		typ := l.detectType(b)
		if typ != l.typ || typ.canNest() {
			t := l.token()
			if t != nil {
				tokens = append(tokens, t)
			}
			l.typ = typ
			l.current = []byte{}
		}
		l.current = append(l.current, b)
	}

	tokens = append(tokens, l.token())

	return tokens, nil
}

func (l *lexer) readByte() (byte, bool, error) {
	pt := make([]byte, 1)
	n, err := l.r.Read(pt)
	if err != nil {
		return 0, false, err
	}
	if n == 0 {
		return 0, false, nil
	}
	return pt[0], true, nil
}

func (l *lexer) detectType(b byte) byteType {
	if l.typ == typeEscape {
		return typeEscapedString
	}

	switch b {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return typeNumber
	case '"':
		return typeQuote
	case '(':
		return typeListStart
	case ')':
		return typeListEnd
	case ' ':
		return typeListSeparator
	case '\\':
		return typeEscape
	default:
		return typeString
	}
}

func (l *lexer) token() Token {
	switch l.typ {
	case typeNumber:
		return &NumberToken{
			numbers: l.current,
		}
	case typeQuote:
		return &QuoteToken{}
	case typeString:
		return &StringToken{
			chars: l.current,
		}
	case typeListStart:
		return &ListStartToken{}
	case typeListEnd:
		return &ListEndToken{}
	case typeListSeparator:
		return &ListSeparatorToken{}
	case typeEscape:
		return &EscapeToken{}
	case typeEscapedString:
		return &EscapedStringToken{
			chr: l.current[0],
		}
	default:
		return nil
	}
}
