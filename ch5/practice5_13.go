package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var crawled map[string]bool

func main() {
	breathFirst2(crawl2, os.Args[1:])
}

func breathFirst2(f func(item string) []string, workList []string) {
	seen := make(map[string]bool)
	for len(workList) > 0 {
		items := workList
		workList = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				workList = append(workList, f(item)...)
			}
		}
	}
}

func crawl2(url string) []string {
	normalizeUrl := func(url string) string {
		url = strings.Split(url, "#")[0]
		if !strings.HasSuffix(url, "/") {
			url += "/"
		}
		return url
	}

	url = normalizeUrl(url)

	var domain string
	if domain == "" {
		split := strings.Split(url, "/")
		domain = split[0] + "//" + split[2]
	}

	// findlinks
	fmt.Println(url)
	list, err := Extract3(url)
	if err != nil {
		log.Print(err)
	}

	// set crawled
	if len(crawled) == 0 {
		crawled = map[string]bool{url: true}
	} else {
		crawled[url] = true
	}

	// save file
	dir := url[8 : len(url)-1]
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		resp, _ := http.Get(url)

		if err := os.MkdirAll(dir, 0775); err != nil {
			fmt.Fprintf(os.Stderr, "cannot create directory")
			return nil
		}
		filePath := dir + "/index.html"
		file, _ := os.Create(filePath)
		io.Copy(file, resp.Body)
		resp.Body.Close()
		file.Close()
	}

	// next
	var result []string
	for _, listItem := range list {
		listItem = normalizeUrl(listItem)
		if strings.HasPrefix(listItem, domain) && !crawled[listItem] {
			result = append(result, listItem)
		}
	}
	return result
}

func Extract3(url string) ([]string, error) {
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

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode6(doc, visitNode, nil)
	return links, nil
}

func forEachNode6(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode6(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
