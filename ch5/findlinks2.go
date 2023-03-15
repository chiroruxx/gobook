package main

import (
	"fmt"
	"gobook/ch5/findlinks2"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findlinks2.FindLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}
