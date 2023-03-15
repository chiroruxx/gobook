package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}

	return true
}

func main() {
	text := os.Args[1:]

	buffer := new(strings.Builder)
	buffer.WriteString("\"" + strings.Join(text, " ") + "\" is ")

	if !IsPalindrome(sort.StringSlice(text)) {
		buffer.WriteString("not ")
	}

	buffer.WriteString("a palindrome.")
	fmt.Println(buffer.String())
}
