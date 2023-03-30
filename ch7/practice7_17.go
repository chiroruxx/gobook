package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var args []xml.StartElement
	for _, arg := range os.Args[1:] {
		args = append(args, buildToken(arg))
	}

	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAllToken(stack, args) {
				var names []string
				for _, item := range stack {
					names = append(names, toString(item))
				}

				fmt.Printf("%s: %s\n", strings.Join(names, " "), tok)
			}
		}
	}
}

func buildToken(text string) (token xml.StartElement) {
	type state int
	const (
		initialState state = iota
		classState
		idState
	)
	var buffer []byte
	now := initialState

	action := func() {
		switch now {
		case initialState:
			name := xml.Name{
				Local: string(buffer),
			}
			token.Name = name
		case classState:
			name := xml.Name{
				Local: "class",
			}
			attr := xml.Attr{
				Name:  name,
				Value: string(buffer),
			}
			token.Attr = append(token.Attr, attr)
		case idState:
			name := xml.Name{
				Local: "id",
			}
			attr := xml.Attr{
				Name:  name,
				Value: string(buffer),
			}
			token.Attr = append(token.Attr, attr)
		}
		var nb []byte
		buffer = nb
	}

	bytes := []byte(text)
	for _, b := range bytes {
		switch b {
		case '.':
			action()
			now = classState
		case '#':
			action()
			now = idState
		default:
			buffer = append(buffer, b)
		}
	}
	action()

	return token
}

func containsAllToken(x, y []xml.StartElement) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if isSameToken(y[0], x[0]) {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func isSameToken(needle, haystack xml.StartElement) bool {
	if needle.Name.Local != haystack.Name.Local {
		return false
	}
	for _, item := range needle.Attr {
		if !hasAttr(item, haystack) {
			return false
		}
	}
	return true
}

func hasAttr(needle xml.Attr, haystack xml.StartElement) bool {
	for _, item := range haystack.Attr {
		if needle.Name.Local == item.Name.Local && needle.Value == item.Value {
			return true
		}
	}
	return false
}

func toString(tok xml.StartElement) string {
	result := tok.Name.Local
	for _, attr := range tok.Attr {
		switch attr.Name.Local {
		case "id":
			result += "#" + attr.Value
		case "class":
			result += "." + attr.Value
		}
	}

	return result
}
