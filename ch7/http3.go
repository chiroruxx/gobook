package main

import (
	"gobook/ch7/http3"
	"log"
	"net/http"
)

func main() {
	db := http3.Database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.List))
	mux.Handle("/price", http.HandlerFunc(db.Price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
