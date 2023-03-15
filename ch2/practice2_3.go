package main

import (
	"fmt"
	"gobook/ch2/practice2_3"
	"os"
	"strconv"
)

func main() {
	value, _ := strconv.ParseUint(os.Args[1], 10, 64)
	count := practice2_3.PopCount(value)

	fmt.Println(count)
}
