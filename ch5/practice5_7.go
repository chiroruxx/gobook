package main

import (
	"gobook/ch5/outline2"
	"gobook/ch5/practice5_7"
	"log"
	"os"
)

func main() {
	node, err := outline2.GetNode(os.Args[1])
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	outline2.ForEachNode(node, practice5_7.StartElement, practice5_7.EndElement)
}
