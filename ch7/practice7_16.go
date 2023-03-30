package main

import (
	"gobook/ch7/practice7_16"
	"log"
	"net/http"
)

func main() {
	server := practice7_16.NewServer()
	log.Fatal(http.ListenAndServe("localhost:8000", server))
}
