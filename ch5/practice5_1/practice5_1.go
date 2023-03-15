package practice5_1

import (
	"golang.org/x/net/html"
)

func Visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	if n.FirstChild != nil {
		links = Visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = Visit(links, n.NextSibling)
	}

	return links
}
