package main

import (
	"fmt"
	"gobook/ch6/practice6_1"
)

func main() {
	var x practice6_1.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y := x.Copy()
	x.Clear()

	fmt.Println(x.String(), y.String())
}
