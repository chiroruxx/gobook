package main

import (
	"fmt"
	"gobook/ch2/popcount"
	"os"
	"strconv"
)

func main() {
	value, _ := strconv.ParseUint(os.Args[1], 10, 64)
	count := popcount.PopCount(value)

	fmt.Println(count)
}
