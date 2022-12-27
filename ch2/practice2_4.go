// package popcount
package main

import (
	"fmt"
	"os"
	"strconv"
)

func PopCount(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		value := x >> i
		if value%2 == 1 {
			count++
		}
	}

	return count
}

func main() {
	value, _ := strconv.ParseUint(os.Args[1], 10, 64)
	count := PopCount(value)

	fmt.Println(count)
}
