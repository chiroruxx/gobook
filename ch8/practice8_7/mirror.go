package practice8_7

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const workerCount = 10
const limit = 100

func Mirror(urlString string) error {
	u, err := newURL(urlString, nil)
	if err != nil {
		return err
	}

	foundLinks := make(chan []*URL, workerCount)
	shouldMirror := make(chan *URL, limit)

	go func(u *URL) {
		foundLinks <- []*URL{u}
	}(u)

	for count := 0; count < workerCount; count++ {
		go func() {
			for {
				u := <-shouldMirror
				founds, err := mirror(*u)
				if err != nil {
					log.Fatal(err)
				}

				foundLinks <- founds
			}
		}()
	}

	mirrored := make(map[string]bool)
	var reachLimit bool
	var limitCount int

	for links := range foundLinks {
		for _, link := range links {
			if limitCount >= limit {
				reachLimit = true
				break
			}

			if _, ok := mirrored[link.String()]; ok {
				continue
			}

			mirrored[link.String()] = true
			shouldMirror <- link
			limitCount++
		}
		if reachLimit {
			break
		}
	}

	return nil
}

func mirror(u URL) (foundLinks []*URL, err error) {
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "got status %d in %s\n", resp.StatusCode, u)
		return []*URL{}, nil
	}

	content, foundLinks, err := replace(resp.Body, u)
	if err != nil {
		return nil, err
	}

	if err := writeFile(content, u); err != nil {
		return nil, err
	}

	return foundLinks, nil
}
