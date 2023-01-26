package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	c1 := sha256.Sum256([]byte(os.Args[1]))
	c2 := sha256.Sum256([]byte(os.Args[2]))
	var count int

	for key, value := range c1 {
		if value != c2[key] {
			count++
		}
	}
	fmt.Print(count)
}
