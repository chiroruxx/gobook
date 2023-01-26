package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var option = flag.String("a", "256", "algorithm")

func main() {
	flag.Parse()

	input := bufio.NewScanner(os.Stdin)
	var plain string
	for input.Scan() {
		plain += input.Text()
	}

	switch *option {
	case "256":
		fmt.Printf("%x", sha256.Sum256([]byte(plain)))
	case "384":
		fmt.Printf("%x", sha512.Sum384([]byte(plain)))
	case "512":
		fmt.Printf("%x", sha512.Sum512([]byte(plain)))
	}
}
