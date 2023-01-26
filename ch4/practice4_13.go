package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

type movie struct {
	Title  string
	Poster string
}

var apiKey2 = flag.String("key", "", "your API Key")
var title2 = flag.String("title", "", "movie title")

func main() {
	flag.Parse()

	if *title2 == "" {
		fmt.Fprintf(os.Stderr, "poster: you should set movie title.\n")
		os.Exit(1)
	}

	if *apiKey2 == "" {
		fmt.Fprintf(os.Stderr, "poster: you should set your API Key.\n")
		os.Exit(1)
	}

	resp, err := http.Get("https://www.omdbapi.com/?apikey=" + *apiKey2 + "&t=" + *title2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "poster: %v\n", err)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "poster: %s\n", resp.Status)
		resp.Body.Close()
		os.Exit(1)
	}

	var movie movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "poster: %v\n", err)
		os.Exit(1)
	}
	resp.Body.Close()

	resp, err = http.Get(movie.Poster)
	if err != nil {
		fmt.Fprintf(os.Stderr, "poster: %v\n", err)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "poster: %s\n", resp.Status)
		resp.Body.Close()
		os.Exit(1)
	}

	fileType := path.Ext(movie.Poster)
	poster, err := os.Create(movie.Title + fileType)
	if err != nil {
		poster.Close()
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "poster: %v\n", err)
		os.Exit(1)
	}
	_, err = io.Copy(poster, resp.Body)
	poster.Close()
	resp.Body.Close()
}
