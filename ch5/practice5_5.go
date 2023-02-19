package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "countWordsAndImages: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("words: %d\nimages: %d\n", words, images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	} else if n.Type == html.TextNode {
		words += countWords(n.Data)
	}

	if n.FirstChild != nil {
		cw, ci := countWordsAndImages(n.FirstChild)
		words += cw
		images += ci
	}
	if n.NextSibling != nil {
		nw, ni := countWordsAndImages(n.NextSibling)
		words += nw
		images += ni
	}
	return
}

func countWords(text string) (count int) {
	text = strings.TrimSpace(text)
	words := strings.Split(text, " ")

	for _, word := range words {
		if word != "" {
			count++
		}
	}

	return
}
