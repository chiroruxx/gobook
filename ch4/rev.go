package main

import (
	"fmt"
	"gobook/ch4/rev"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	rev.Reverse(a[:])
	fmt.Println(a)

	s := []int{0, 1, 2, 3, 4, 5}
	rev.Reverse(s[:2])
	rev.Reverse(s[2:])
	rev.Reverse(s)
	fmt.Println(s)
	fmt.Println(cap(s), len(s))
}
