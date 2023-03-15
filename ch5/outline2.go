package main

import (
	"gobook/ch5/outline2"
	"log"
	"os"
)

func main() {
	node, err := outline2.GetNode(os.Args[1])
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	outline2.ForEachNode(node, outline2.StartElement, outline2.EndElement)
}
