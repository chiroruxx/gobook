package practice5_13

import (
	"fmt"
	"gobook/ch5/links"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var crawled map[string]bool

func Crawl(url string) []string {
	normalizeUrl := func(url string) string {
		url = strings.Split(url, "#")[0]
		if !strings.HasSuffix(url, "/") {
			url += "/"
		}
		return url
	}

	url = normalizeUrl(url)

	var domain string
	if domain == "" {
		split := strings.Split(url, "/")
		domain = split[0] + "//" + split[2]
	}

	// findlinks
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	// set crawled
	if len(crawled) == 0 {
		crawled = map[string]bool{url: true}
	} else {
		crawled[url] = true
	}

	// save file
	dir := url[8 : len(url)-1]
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		resp, _ := http.Get(url)

		if err := os.MkdirAll(dir, 0775); err != nil {
			fmt.Fprintf(os.Stderr, "cannot create directory")
			return nil
		}
		filePath := dir + "/index.html"
		file, _ := os.Create(filePath)
		io.Copy(file, resp.Body)
		resp.Body.Close()
		file.Close()
	}

	// next
	var result []string
	for _, listItem := range list {
		listItem = normalizeUrl(listItem)
		if strings.HasPrefix(listItem, domain) && !crawled[listItem] {
			result = append(result, listItem)
		}
	}
	return result
}
