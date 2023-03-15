package main

import (
	"gobook/ch7/http1"
	"gobook/ch7/practice7_11"
	"log"
	"net/http"
)

func main() {
	db := practice7_11.NewDatabase(map[string]http1.Dollars{"shoes": 50, "socks": 5})
	http.HandleFunc("/", db.ListOrStore)
	http.HandleFunc("/list", db.List)
	http.HandleFunc("/item", db.OperateItem)
	http.HandleFunc("/price", db.Price)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
