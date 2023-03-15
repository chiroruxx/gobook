package main

import (
	"gobook/ch5/findlinks3"
	"gobook/ch5/practice5_13"
	"os"
)

func main() {
	findlinks3.BreathFirst(practice5_13.Crawl, os.Args[1:])
}
