package main

import (
	"fmt"
	"os"
)

func countBytes(s string) map[byte]int {
	result := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		result[s[i]]++
	}

	return result
}

func equals(one, another map[byte]int) bool {
	if len(one) != len(another) {
		return false
	}

	for key, value := range one {
		if another[key] != value {
			return false
		}
	}

	return true
}

func main() {
	one := os.Args[1]
	another := os.Args[2]

	count1 := countBytes(one)
	count2 := countBytes(another)

	if equals(count1, count2) {
		fmt.Printf("%s is an anagram of %s\n", one, another)
	} else {
		fmt.Printf("%s is not an anagram of %s\n", one, another)
	}
}
