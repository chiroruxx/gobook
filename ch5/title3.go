package main

import (
	"fmt"
	"gobook/ch5/title3"
	"log"
)

func main() {
	title, err := title3.Title("https://go.dev/")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(title)
}
