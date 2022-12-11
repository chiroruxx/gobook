package main

import (
	"fmt"
	"os"
)

func main() {
	for index, value := range os.Args[1:] {
		fmt.Print(index)
		fmt.Println(":" + value)
	}
}
