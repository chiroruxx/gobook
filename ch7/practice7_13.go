package main

import (
	"fmt"
	"gobook/ch7/practice7_13"
	"log"
)

func main() {
	input := "sin(-x)*pow(1.5,-r)"

	expr, err := practice7_13.ParseAndCheck(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("in:  %s\n", input)
	fmt.Printf("out: %s", expr)
}
