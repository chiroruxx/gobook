package practice8_7

import (
	"bytes"
	"golang.org/x/net/html"
	"io"
)

func replace(content io.Reader, requestUrl URL) (replaced []byte, foundLinks []*URL, err error) {
	root, err := html.Parse(content)
	if err != nil {
		return nil, nil, err
	}

	foundLinks, err = eachNode(root, requestUrl)
	if err != nil {
		return nil, nil, err
	}

	var buffer bytes.Buffer
	if err = html.Render(&buffer, root); err != nil {
		return nil, nil, err
	}

	return buffer.Bytes(), foundLinks, err
}

func eachNode(node *html.Node, requestUrl URL) (foundLinks []*URL, err error) {
	foundLinks, err = replaceNode(node, requestUrl)
	if err != nil {
		return nil, err
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		childLinks, err := eachNode(child, requestUrl)
		if err != nil {
			return nil, err
		}
		foundLinks = append(foundLinks, childLinks...)
	}

	return foundLinks, nil
}

func replaceNode(node *html.Node, requestUrl URL) (foundLinks []*URL, err error) {
	if node.Type != html.ElementNode || node.Data != "a" {
		return nil, nil
	}

	for attrNo, attribute := range node.Attr {
		if attribute.Key != "href" {
			continue
		}

		u, err := newURL(attribute.Val, &requestUrl)
		if err != nil {
			return nil, err
		}

		if !u.IsSameHost(requestUrl) {
			continue
		}

		replaced := fileLink(*u)
		node.Attr[attrNo].Val = replaced

		foundLinks = append(foundLinks, u)
	}

	return foundLinks, nil
}
