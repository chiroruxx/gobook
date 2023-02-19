package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

var depth2 int

func main() {
	node, err := getNode2(os.Args[1])
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	forEachNode2(node, startElement2, endElement2)
}

func getNode2(url string) (*html.Node, error) {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return doc, nil
}

func forEachNode2(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode2(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func startElement2(n *html.Node) {
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
		fmt.Printf("%*s%s\n", depth2*2, "", content)
	}

	if n.Type == html.ElementNode {
		depth2++
	}
}

func endElement2(n *html.Node) {
	if n.Type == html.ElementNode {
		depth2--
	}

	var tag string
	if n.Type == html.ElementNode {
		tag = buildElementEndTag(n)
	}

	if tag != "" {
		fmt.Printf("%*s%s\n", depth2*2, "", tag)
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
