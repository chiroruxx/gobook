package title1

import (
	"fmt"
	"gobook/ch5/outline2"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func Title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if ct := resp.Header.Get("Content-Type"); ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	outline2.ForEachNode(doc, visitNode, nil)
	return nil
}
