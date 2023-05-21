package main

import (
	"gobook/ch8/crawl2"
	"os"
)

func main() {
	workList := make(chan []string)
	var n int

	n++
	go func() {
		workList <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-workList
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					workList <- crawl2.Crawl(link)
				}(link)
			}
		}
	}
}
