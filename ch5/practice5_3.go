package main

import (
	"fmt"
	"gobook/ch5/practice5_3"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, text := range practice5_3.Visit(nil, doc) {
		fmt.Println(text)
	}
}
