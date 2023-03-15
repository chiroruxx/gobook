package main

import (
	"fmt"
	"gobook/ch5/outline2"
	"gobook/ch5/practice5_8"
	"log"
	"os"
)

func main() {
	node, err := outline2.GetNode(os.Args[1])
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	node = practice5_8.ElementById(node, os.Args[2])
	if node == nil {
		fmt.Println("No node found.")
	} else {
		fmt.Println(node.Data)
	}
}
