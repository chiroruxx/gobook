package main

import (
	"fmt"
	"gobook/ch4/practice4_3"
)

func main() {
	array := [4]int{0, 1, 2, 3}
	practice4_3.Reverse(&array)
	fmt.Print(array)
}
