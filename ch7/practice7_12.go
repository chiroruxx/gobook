package main

import (
	"gobook/ch7/practice7_12"
	"log"
	"net/http"
)

func main() {
	db := practice7_12.Database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.List))
	mux.Handle("/price", http.HandlerFunc(db.Price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
