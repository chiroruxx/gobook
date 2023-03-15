package main

import (
	"fmt"
	"gobook/ch6/practice6_5"
)

func main() {
	var x practice6_5.IntSet
	x.AddAll(1, 144, 9, 12, 145)
	fmt.Println(x.String())
}
