package crawl2

import (
	"fmt"
	"gobook/ch5/links"
	"log"
)

var tokens = make(chan struct{}, 20)

func Crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}
