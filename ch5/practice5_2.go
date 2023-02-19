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
	for name, count := range mapping(make(map[string]int), doc) {
		fmt.Printf("%s:\t%d\n", name, count)
	}
}

func mapping(tagMap map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		tagMap[n.Data]++
	}

	if n.FirstChild != nil {
		tagMap = mapping(tagMap, n.FirstChild)
	}
	if n.NextSibling != nil {
		tagMap = mapping(tagMap, n.NextSibling)
	}

	return tagMap
}
