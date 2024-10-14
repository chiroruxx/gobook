package practice12_10

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"text/scanner"
)

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func (lex *lexer) next() {
	lex.token = lex.scan.Scan()
}

func (lex *lexer) text() string {
	return lex.scan.TokenText()
}

func (lex *lexer) consume(want rune) {
	if lex.token != want {
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

type Token any

type Symbol struct {
	Value string
}

type String struct {
	Value string
}

type Int struct {
	Value int
}

type StartList struct {
}

type EndList struct {
}

func token(lex *lexer) (Token, error) {
	switch lex.token {
	case scanner.Ident:
		return Symbol{
			Value: lex.text(),
		}, nil
	case scanner.String:
		s, _ := strconv.Unquote(lex.text())
		return String{
			Value: s,
		}, nil
	case scanner.Int:
		fmt.Println("int")
		i, _ := strconv.Atoi(lex.text())
		return Int{
			Value: i,
		}, nil
	case '(':
		return StartList{}, nil
	case ')':
		return EndList{}, nil
	case scanner.EOF:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected token %q", lex.text())
	}
}

func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		switch lex.text() {
		case "nil":
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		case "t":
			v.SetBool(true)
			lex.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text())
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text())
		v.SetInt(int64(i))
		lex.next()
		return
	case scanner.Float:
		f, _ := strconv.ParseFloat(lex.text(), 64)
		v.SetFloat(f)
		lex.next()
		return
	case '(':
		lex.next() // '('
		readList(lex, v)
		lex.next() // ')'
		return
	default:
		panic(fmt.Sprintf("unexpected token %q", lex.text()))
	}
}

func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array:
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}
	case reflect.Slice:
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}
	case reflect.Struct:
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}
	case reflect.Map:
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Key()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}
	case reflect.Interface:
		for !endList(lex) {
			typ := reflect.ValueOf(lex.text())
			typ.Kind()
			rTyp, err := getReflectTypeFromString(typ.String())
			if err != nil {
				panic(err)
			}

			value := reflect.New(rTyp).Elem()
			lex.next()
			read(lex, value)
			v.Set(value)
		}
	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	default:
		return false
	}
}

func Unmarshal(data []byte, out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader(data))
	lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}

func getReflectTypeFromString(s string) (reflect.Type, error) {
	type typeString string
	const (
		tsString  = "string"
		tsInt     = "int"
		tsFloat64 = "float64"
		tsBool    = "bool"
	)

	ts := typeString(s)
	switch ts {
	case tsString:
		return reflect.TypeOf(""), nil
	case tsInt:
		return reflect.TypeOf(0), nil
	case tsFloat64:
		return reflect.TypeOf(float64(0)), nil
	case tsBool:
		return reflect.TypeOf(false), nil
	}

	return nil, fmt.Errorf("unknown type string %q", s)
}
