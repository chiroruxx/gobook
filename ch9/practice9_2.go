package main

import (
	"fmt"
	"os"
	"strconv"

	"gobook/ch9/practice9_2"
)

func main() {
	value, _ := strconv.ParseUint(os.Args[1], 10, 64)
	count := practice9_2.PopCount(value)

	fmt.Println(count)
}
