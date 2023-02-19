package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, text := range visit3(nil, doc) {
		fmt.Println(text)
	}
}

func visit3(texts []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text != "" {
			texts = append(texts, text)
		}
	}

	if n.Data != "script" && n.Data != "style" {
		if n.FirstChild != nil {
			texts = visit3(texts, n.FirstChild)
		}
	}
	if n.NextSibling != nil {
		texts = visit3(texts, n.NextSibling)
	}

	return texts
}
