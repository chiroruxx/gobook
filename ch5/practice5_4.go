package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit4(nil, doc) {
		fmt.Println(link)
	}
}

func visit4(links []string, n *html.Node) []string {
	link := getLink(n)
	if link != nil {
		links = append(links, *link)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit4(links, c)
	}

	return links
}

func getLink(n *html.Node) *string {
	if n.Type != html.ElementNode {
		return nil
	}

	key := getLinkAttributeKey(n)

	if key == "" {
		return nil
	}

	for _, attribute := range n.Attr {
		if attribute.Key == key {
			return &attribute.Val
		}
	}

	return nil
}

func getLinkAttributeKey(n *html.Node) (key string) {
	switch n.Data {
	case "a", "link":
		key = "href"
	case "img", "script":
		key = "src"
	}
	return
}
