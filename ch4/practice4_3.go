package main

import "fmt"

func reverse2(s *[4]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	array := [4]int{0, 1, 2, 3}
	reverse2(&array)
	fmt.Print(array)
}
