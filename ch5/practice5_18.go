package main

import (
	"fmt"
	"gobook/ch5/practice5_18"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		name, _, err := practice5_18.Fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
		fmt.Println(name)
	}
}
