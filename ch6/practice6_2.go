package main

import (
	"fmt"
	"gobook/ch6/practice6_2"
)

func main() {
	var x practice6_2.IntSet
	x.AddAll(1, 144, 9)
	fmt.Println(x.String())
}
