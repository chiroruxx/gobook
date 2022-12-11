package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	inputs := strings.Split(strings.Repeat("x", 100_000), "")

	start := time.Now()

	var s, sep string
	for i := 1; i < len(inputs); i++ {
		s += sep + inputs[i]
		sep = " "
	}
	//fmt.Println(s)

	secs := time.Since(start).Seconds()
	fmt.Printf("%.3fs\n", secs)

	start = time.Now()

	strings.Join(inputs[1:], " ")
	//fmt.Println(strings.Join(inputs[1:], " "))

	secs = time.Since(start).Seconds()
	fmt.Printf("%.3fs", secs)
}
