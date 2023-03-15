package main

import (
	"fmt"
	"gobook/ch7/practice7_3"
)

func main() {
	t := new(practice7_3.Tree)
	fmt.Println(t)
	t.Add(5)
	fmt.Println(t)
	t.Add(0)
	fmt.Println(t)
	t.Add(1)
	fmt.Println(t)
	t.Add(4)
	fmt.Println(t)
	t.Add(3)
	fmt.Println(t)
}
