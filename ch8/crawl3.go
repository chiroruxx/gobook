package main

import (
	"gobook/ch8/crawl1"
	"os"
)

func main() {
	workList := make(chan []string)
	unseenLinks := make(chan string)

	go func() {
		workList <- os.Args[1:]
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl1.Crawl(link)
				go func() {
					workList <- foundLinks
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
