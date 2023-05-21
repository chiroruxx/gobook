package main

import (
	"fmt"
	"gobook/ch8/practice8_7"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "practice8_7: url does not given")
		fmt.Fprintln(os.Stderr, "usage: go run practice8_7 [url]")
		os.Exit(1)
	}

	url := os.Args[1]

	if err := practice8_7.Mirror(url); err != nil {
		log.Fatal(err)
	}
}
