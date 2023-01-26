package main

import "fmt"

func rotate(s []int, i int) []int {
	return append(s[i:], s[:i]...)
}

func main() {
	array := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(rotate(array, 2))
}
