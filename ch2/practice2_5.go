package main

import (
	"fmt"
	"gobook/ch2/practice2_5"
	"os"
	"strconv"
)

func main() {
	value, _ := strconv.ParseUint(os.Args[1], 10, 64)
	count := practice2_5.PopCount(value)

	fmt.Println(count)
}
