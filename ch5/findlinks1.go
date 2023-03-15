package main

import (
	"fmt"
	"gobook/ch5/findlinks1"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range findlinks1.Visit(nil, doc) {
		fmt.Println(link)
	}
}
