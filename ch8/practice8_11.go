package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

type result struct {
	name    string
	content io.ReadCloser
}

var done = make(chan struct{})

func main() {
	r := mirroredQuery(os.Args[1:])
	if err := createFile(r); err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	}

	fmt.Println(r.name)
}

func mirroredQuery(urls []string) *result {
	responses := make(chan *result, len(urls))

	for _, url := range urls {
		url := url
		go func() {
			r, err := fetch(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
				return
			}
			fmt.Println("Done: " + url)
			close(done)
			responses <- r
		}()
	}

	r := <-responses
	close(responses)
	return r
}

func canceled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func fetch(url string) (*result, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Cancel = done

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	res := &result{
		name:    resp.Request.URL.Path,
		content: resp.Body,
	}
	return res, nil
}

func createFile(r *result) error {
	defer r.content.Close()

	local := path.Base(r.name)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return err
	}

	if _, err := io.Copy(f, r.content); err != nil {
		return err
	}

	if err := f.Close(); err == nil {
		return err
	}

	return nil
}
