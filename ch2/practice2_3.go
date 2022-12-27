// package popcount
package main

import (
	"fmt"
	"os"
	"strconv"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var value byte
	for i := 0; i < 8; i++ {
		value += pc[byte(x>>(i*8))]
	}
	return int(value)
}

func main() {
	value, _ := strconv.ParseUint(os.Args[1], 10, 64)
	count := PopCount(value)

	fmt.Println(count)
}
