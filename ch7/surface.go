package main

import (
	"gobook/ch7/surface"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", surface.Plot)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
