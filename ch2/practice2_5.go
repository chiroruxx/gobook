// package popcount
package main

import (
	"fmt"
	"os"
	"strconv"
)

func PopCount(x uint64) int {
	count := 0
	for x != 0 {
		count++
		x = x & (x - 1)
	}

	return count
}

func main() {
	value, _ := strconv.ParseUint(os.Args[1], 10, 64)
	count := PopCount(value)

	fmt.Println(count)
}
