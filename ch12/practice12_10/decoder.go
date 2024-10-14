package practice12_10

import (
	"fmt"
	"io"
	"reflect"
	"text/scanner"
)

type Decoder struct {
	r   io.Reader
	lex *lexer
}

func NewDecoder(r io.Reader) *Decoder {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(r)

	return &Decoder{
		r:   r,
		lex: lex,
	}
}

func (d *Decoder) Decode(v any) (err error) {
	d.lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", d.lex.scan.Position, x)
		}
	}()
	read(d.lex, reflect.ValueOf(v).Elem())
	return nil
}

func (d *Decoder) Token() (t Token, err error) {
	d.lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", d.lex.scan.Position, x)
		}
	}()

	return token(d.lex)
}
