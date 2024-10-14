package lexer

import (
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_Lexer_readByte(t *testing.T) {
	l := lexer{
		r: strings.NewReader("abc"),
	}

	type res struct {
		b   byte
		ok  bool
		err error
	}

	expects := []res{
		{b: 'a', ok: true, err: nil},
		{b: 'b', ok: true, err: nil},
		{b: 'c', ok: true, err: nil},
		{ok: false, err: io.EOF},
	}

	for i, expected := range expects {
		b, ok, err := l.readByte()
		if !errors.Is(err, expected.err) {
			t.Errorf("readByte() %d error = %v, wantErr %v", i, err, expected.err)
			return
		}
		if b != expected.b {
			t.Errorf("readByte() %d b = %v, want %v", i, b, expected.b)
		}
		if ok != expected.ok {
			t.Errorf("readByte() %d ok = %v, want %v", i, ok, expected.ok)
		}
	}
}

func TestLexer_lex(t *testing.T) {
	type fields struct {
		r         io.Reader
		proceeded []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Token
		wantErr bool
	}{
		{
			"number",
			fields{
				r: strings.NewReader("4"),
			},
			[]Token{
				&NumberToken{
					numbers: []byte{'4'},
				},
			},
			false,
		},
		{
			"2 digit number",
			fields{
				r: strings.NewReader("42"),
			},
			[]Token{
				&NumberToken{
					numbers: []byte{'4', '2'},
				},
			},
			false,
		},
		{
			"quote",
			fields{
				r: strings.NewReader(`"`),
			},
			[]Token{
				&QuoteToken{},
			},
			false,
		},
		{
			"string",
			fields{
				r: strings.NewReader(`"a"`),
			},
			[]Token{
				&QuoteToken{},
				&StringToken{
					chars: []byte{'a'},
				},
				&QuoteToken{},
			},
			false,
		},
		{
			"list",
			fields{
				r: strings.NewReader("(1 2 3)"),
			},
			[]Token{
				&ListStartToken{},
				&NumberToken{
					numbers: []byte{'1'},
				},
				&ListSeparatorToken{},
				&NumberToken{
					numbers: []byte{'2'},
				},
				&ListSeparatorToken{},
				&NumberToken{
					numbers: []byte{'3'},
				},
				&ListEndToken{},
			},
			false,
		},
		{
			"escape",
			fields{
				r: strings.NewReader(`\"`),
			},
			[]Token{
				&EscapeToken{},
				&EscapedStringToken{
					chr: '"',
				},
			},
			false,
		},
		{
			"sample",
			fields{
				r: strings.NewReader(`((Name "John") (Age 18))`),
			},
			[]Token{
				&ListStartToken{},
				&ListStartToken{},
				&StringToken{
					chars: []byte("Name"),
				},
				&ListSeparatorToken{},
				&QuoteToken{},
				&StringToken{
					chars: []byte("John"),
				},
				&QuoteToken{},
				&ListEndToken{},
				&ListSeparatorToken{},
				&ListStartToken{},
				&StringToken{
					chars: []byte("Age"),
				},
				&ListSeparatorToken{},
				&NumberToken{
					numbers: []byte("18"),
				},
				&ListEndToken{},
				&ListEndToken{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &lexer{
				r: tt.fields.r,
			}
			got, err := l.Lex()
			if (err != nil) != tt.wantErr {
				t.Errorf("Lex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lex() got = %v, want %v", got, tt.want)
			}
		})
	}
}
