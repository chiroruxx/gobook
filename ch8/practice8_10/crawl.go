package practice8_10

import (
	"fmt"
	"log"
	"os"
)

var done chan struct{}

func Crawl(url string) []string {
	if Canceled() {
		return []string{}
	}

	fmt.Println(url)
	list, err := extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func CancelFunc() func() {
	return func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}
}

func Canceled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
