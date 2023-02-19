package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

func ElementsByTagName(doc *html.Node, names ...string) []*html.Node {
	var founds []*html.Node

	if isElementByNames(doc, names...) {
		founds = append(founds, doc)
	}

	if child := doc.FirstChild; child != nil {
		founds = append(founds, ElementsByTagName(child, names...)...)
	}

	if next := doc.NextSibling; next != nil {
		founds = append(founds, ElementsByTagName(next, names...)...)
	}

	return founds
}

func isElementByName(doc *html.Node, name string) bool {
	return doc.Type == html.ElementNode && doc.Data == name
}

func isElementByNames(doc *html.Node, names ...string) bool {
	result := false

	for _, name := range names {
		result = result || isElementByName(doc, name)
	}

	return result
}

func getDoc(url string) (*html.Node, error) {
	resp, err := http.Get(url)
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

func nodeToString(node *html.Node) string {
	text := "<" + node.Data

	for _, attribute := range node.Attr {
		text += " " + attribute.Key + "=" + attribute.Val
	}

	text += ">"

	return text
}

func main() {
	url := "https://go.dev/"
	doc, err := getDoc(url)
	if err != nil {
		log.Fatal(err)
		return
	}

	nodes := ElementsByTagName(doc, os.Args[1:]...)

	for _, node := range nodes {
		fmt.Println(nodeToString(node))
	}
}
