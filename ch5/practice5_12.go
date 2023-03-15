package main

import (
	"fmt"
	"gobook/ch5/outline2"
	"golang.org/x/net/html"
	"log"
	"os"
)

func main() {
	node, err := outline2.GetNode(os.Args[1])
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	var depth int
	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}
	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}

	outline2.ForEachNode(node, startElement, endElement)
}
