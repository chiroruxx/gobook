package main

import "fmt"

func removeDuplicate(s []string) []string {
	i := 0
	var previous string

	for index, value := range s {
		if index == 0 || previous != value {
			s[i] = value
			i++
		}

		previous = value
	}

	return s[:i]
}

func main() {
	array := []string{"b", "b", "b", "c", "b"}
	fmt.Println(removeDuplicate(array))
}
