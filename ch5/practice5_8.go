package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

var targetId string

func main() {
	node, err := getNode3(os.Args[1])
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	node = ElementById(node, os.Args[2])
	if node == nil {
		fmt.Println("No node found.")
	} else {
		fmt.Println(node.Data)
	}
}

func getNode3(url string) (*html.Node, error) {
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

func ElementById(n *html.Node, id string) *html.Node {
	targetId = id
	return forEachNode3(n, doesntHaveId, nil)
}

func forEachNode3(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil && !pre(n) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode3(c, pre, post)
		if node != nil {
			return node
		}
	}

	if post != nil && !post(n) {
		return n
	}

	return nil
}

func hasId(n *html.Node) bool {
	for _, attribute := range n.Attr {
		if attribute.Key == "id" && attribute.Val == targetId {
			return true
		}
	}
	return false
}

func doesntHaveId(n *html.Node) bool {
	return !hasId(n)
}
