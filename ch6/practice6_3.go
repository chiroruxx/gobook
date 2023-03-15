package main

import (
	"fmt"
	"gobook/ch6/practice6_3"
)

func main() {
	var x, y practice6_3.IntSet
	x.AddAll(1, 144, 9, 12, 145)
	y.AddAll(1, 144, 3, 12, 550)
	x.DifferenceWith(&y)
	fmt.Println(x.String())
}
