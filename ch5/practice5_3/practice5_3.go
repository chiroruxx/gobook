package practice5_3

import (
	"golang.org/x/net/html"
	"strings"
)

func Visit(texts []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text != "" {
			texts = append(texts, text)
		}
	}

	if n.Data != "script" && n.Data != "style" {
		if n.FirstChild != nil {
			texts = Visit(texts, n.FirstChild)
		}
	}
	if n.NextSibling != nil {
		texts = Visit(texts, n.NextSibling)
	}

	return texts
}
