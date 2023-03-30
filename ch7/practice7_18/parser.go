package practice7_18

import (
	"bytes"
	"encoding/xml"
	"regexp"
	"strings"
)

type Node interface {
	String() string
}

type CharData string

var indent = 0

func (d CharData) String() string {
	str := strings.ReplaceAll(string(d), "\n", " ")

	regex := regexp.MustCompile(" +")
	str = regex.ReplaceAllString(str, " ")
	if len(str) > 100 {
		str = str[:100] + "..."
	}

	return "Char<" + str + ">"
}

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e Element) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Elem<")
	buffer.WriteString(e.Type.Local)
	for _, attr := range e.Attr {
		buffer.WriteByte(' ')
		buffer.WriteString(attr.Name.Local)
		buffer.WriteString("=\"")
		buffer.WriteString(attr.Value)
		buffer.WriteByte('"')
	}
	buffer.WriteByte('>')

	indent++
	for _, child := range e.Children {
		buffer.WriteByte('\n')
		buffer.WriteString(strings.Repeat(" ", indent*2))
		buffer.WriteString(child.String())
	}
	indent--

	return buffer.String()
}

type Parser struct{}

func NewParser() *Parser {
	parser := Parser{}

	return &parser
}

func (p *Parser) Parse(tokens *[]xml.Token, sp *xml.StartElement) Node {
	var children []Node
	start := sp

	for len(*tokens) != 0 {
		token := (*tokens)[0]
		*tokens = (*tokens)[1:]

		switch token := token.(type) {
		case xml.StartElement:
			if start == nil {
				start = &token
				continue
			}
		}

		if start == nil {
			panic("start not found.")
		}

		switch token := token.(type) {
		case xml.StartElement:
			children = append(children, p.Parse(tokens, &token))
		case xml.EndElement:
			return Element{
				Type:     start.Name,
				Attr:     start.Attr,
				Children: children,
			}
		case xml.CharData:
			children = append(children, CharData(token))
		}
	}

	panic("loop end")
}
