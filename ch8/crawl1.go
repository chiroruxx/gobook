package main

import (
	"gobook/ch8/crawl1"
	"os"
)

func main() {
	workList := make(chan []string)

	go func() {
		workList <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					workList <- crawl1.Crawl(link)
				}(link)
			}
		}
	}
}
