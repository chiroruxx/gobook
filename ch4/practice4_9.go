package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := make(map[string]int)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		text := in.Text()
		words[text]++
	}

	for text, count := range words {
		fmt.Printf("%s\t %d\n", text, count)
	}
}
