package practice8_6

import (
	"gobook/ch5/links"
	"log"
)

var tokens = make(chan struct{}, 20)

func Crawl(url string) []string {
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}
