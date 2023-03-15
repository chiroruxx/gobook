package practice5_8

import (
	"golang.org/x/net/html"
)

var targetId string

func ElementById(n *html.Node, id string) *html.Node {
	targetId = id
	return forEachNode(n, doesntHaveId, nil)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil && !pre(n) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(c, pre, post)
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
