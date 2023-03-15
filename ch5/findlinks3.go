package main

import (
	"gobook/ch5/findlinks3"
	"os"
)

func main() {
	findlinks3.BreathFirst(findlinks3.Crawl, os.Args[1:])
}
