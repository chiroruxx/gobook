package main

import (
	"gobook/ch3/practice3_4"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", practice3_4.Handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
