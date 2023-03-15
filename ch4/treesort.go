package main

import (
	"fmt"
	"gobook/ch4/treesort"
)

func main() {
	ints := []int{0, 5, 3, 2, 1, 4}
	treesort.Sort(ints)
	fmt.Println(ints)
}
