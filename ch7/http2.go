package main

import (
	"gobook/ch7/http2"
	"log"
	"net/http"
)

func main() {
	db := http2.Database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
