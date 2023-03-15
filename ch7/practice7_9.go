package main

import (
	"gobook/ch7/practice7_9"
	"log"
	"net/http"
)

// main

func main() {
	http.HandleFunc("/", practice7_9.Handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
