package practice5_7

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

var depth int

func StartElement(n *html.Node) {
	var content string
	switch n.Type {
	case html.ElementNode:
		content = buildElementStartTag(n)
	case html.DoctypeNode:
		content = buildDocTypeTag(n)
	case html.TextNode:
		content = buildText(n)
	case html.CommentNode:
		content = buildContentTag(n)
	}

	if content != "" {
		fmt.Printf("%*s%s\n", depth*2, "", content)
	}

	if n.Type == html.ElementNode {
		depth++
	}
}

func EndElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
	}

	var tag string
	if n.Type == html.ElementNode {
		tag = buildElementEndTag(n)
	}

	if tag != "" {
		fmt.Printf("%*s%s\n", depth*2, "", tag)
	}
}

func buildDocTypeTag(n *html.Node) string {
	return fmt.Sprintf("<!DOCTYPE %s>\n", n.Data)
}

func buildElementStartTag(n *html.Node) string {
	var buffer bytes.Buffer
	buffer.WriteByte('<')
	buffer.WriteString(n.Data)
	for _, attribute := range n.Attr {
		buffer.WriteString(fmt.Sprintf(" %s=\"%s\"", attribute.Key, attribute.Val))
	}

	if isShortNode(n) {
		buffer.WriteByte('/')
	}

	buffer.WriteByte('>')

	return buffer.String()
}

func buildElementEndTag(n *html.Node) string {
	if isShortNode(n) {
		return ""
	}

	return fmt.Sprintf("</%s>", n.Data)
}

func isShortNode(n *html.Node) bool {
	if n.Data == "script" {
		return false
	}

	return n.FirstChild == nil
}

func buildText(n *html.Node) string {
	text := n.Data
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, "\n", " ")
	return text
}

func buildContentTag(n *html.Node) string {
	return fmt.Sprintf("<!-- %s -->", n.Data)
}
