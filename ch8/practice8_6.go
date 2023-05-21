package main

import (
	"flag"
	"fmt"
	"gobook/ch8/practice8_6"
	"os"
)

var depthLimit = flag.Int("depth", 1, "")

func main() {
	flag.Parse()

	type workListItem struct {
		Urls  []string
		Depth int
	}

	workList := make(chan workListItem)
	var n int

	n++
	go func() {
		workList <- workListItem{
			Urls: os.Args[3:],
		}
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		i := <-workList

		for _, url := range i.Urls {
			if !seen[url] {
				seen[url] = true

				fmt.Println(i.Depth, url)
				if i.Depth >= *depthLimit {
					continue
				}

				n++
				go func(url string, depth int) {
					founds := practice8_6.Crawl(url)
					workList <- workListItem{
						Urls:  founds,
						Depth: depth + 1,
					}
				}(url, i.Depth)
			}
		}
	}
}
