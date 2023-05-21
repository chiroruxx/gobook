package crawl1

import (
	"fmt"
	"gobook/ch5/links"
	"log"
)

func Crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
