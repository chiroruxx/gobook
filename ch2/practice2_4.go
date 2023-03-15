package main

import (
	"fmt"
	"gobook/ch2/practice2_4"
	"os"
	"strconv"
)

func main() {
	value, _ := strconv.ParseUint(os.Args[1], 10, 64)
	count := practice2_4.PopCount(value)

	fmt.Println(count)
}
