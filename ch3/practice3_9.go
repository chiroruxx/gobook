package main

import (
	"gobook/ch3/practice3_9"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", practice3_9.Handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
