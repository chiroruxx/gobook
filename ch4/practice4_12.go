package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type comic struct {
	URL        string `json:"img"`
	Transcript string `json:"transcript"`
}

const indexFilePath = "./index.json"

var indexFlag = flag.Bool("index", false, "create index file")

func main() {
	flag.Parse()

	if *indexFlag {
		createIndexFile()
	} else {
		if len(os.Args) < 2 {
			fmt.Fprintf(os.Stderr, "xkcd: you should set search word!\n")
			os.Exit(1)
		}
		search(os.Args[1])
	}
}

func createIndexFile() {
	if isExistIndexFile() {
		fmt.Fprintf(os.Stderr, "xkcd: index file exists!\n")
		os.Exit(1)
	}
	file, err := os.Create(indexFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "xkcd: %v\n", err)
		os.Exit(1)
	}

	var comics []comic
	for i := 1; i < 100; i++ {
		endpoint := "https://xkcd.com/" + strconv.Itoa(i) + "/info.0.json"
		resp, err := http.Get(endpoint)
		if err != nil {
			fmt.Fprintf(os.Stderr, "xkcd: %v\n", err)
			os.Exit(1)
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			fmt.Fprintf(os.Stderr, "xkcd: %s\n", resp.Status)
		}

		var item comic
		if err = json.NewDecoder(resp.Body).Decode(&item); err != nil {
			fmt.Fprintf(os.Stderr, "xkcd: %v\n", err)
			os.Exit(1)
		}
		comics = append(comics, item)
		time.Sleep(time.Millisecond * 10)
	}

	bytes, err := json.Marshal(comics)
	if err != nil {
		fmt.Fprintf(os.Stderr, "xkcd: %v\n", err)
		os.Exit(1)
	}

	if _, err := file.Write(bytes); err != nil {
		fmt.Fprintf(os.Stderr, "xkcd: %v\n", err)
		os.Exit(1)
	}
	file.Close()

	fmt.Println("created the index file!")
}

func isExistIndexFile() bool {
	_, err := os.Stat(indexFilePath)
	return !os.IsNotExist(err)
}

func search(searchWord string) {
	if !isExistIndexFile() {
		fmt.Fprintf(os.Stderr, "xkcd: index file exists!\n")
		os.Exit(1)
	}
	bytes, err := os.ReadFile(indexFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "xkcd: %v\n", err)
		os.Exit(1)
	}
	var comics []comic
	if err := json.Unmarshal(bytes, &comics); err != nil {
		fmt.Fprintf(os.Stderr, "xkcd: %v\n", err)
		os.Exit(1)
	}

	var result []comic
	for _, comic := range comics {
		if strings.Contains(comic.Transcript, searchWord) {
			result = append(result, comic)
		}
	}

	fmt.Println(strconv.Itoa(len(result)) + "hits.")
	fmt.Println("------")
	for _, comic := range result {
		fmt.Printf("URL: " + comic.URL + "\n")
		fmt.Println("Transcript:")
		fmt.Println(comic.Transcript)
		fmt.Println("------")
	}
}
