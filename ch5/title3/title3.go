package title3

import (
	"fmt"
	"gobook/ch5/outline2"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func Title(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if ct := resp.Header.Get("Content-Type"); ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return "", fmt.Errorf("%s has %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	title, err := soleTitle(doc)
	if err != nil {
		return "", fmt.Errorf("soling %s: %v", url, err)
	}

	return title, nil
}

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
		case bailout{}:
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p)
		}
	}()

	outline2.ForEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}, nil)

	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}
