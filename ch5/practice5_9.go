package main

import (
	"fmt"
	"strings"
)

func main() {
	text := expand("Hello, $Taro How are you?", addSan)
	fmt.Println(text)
}

func expand(s string, f func(string) string) string {
	words := strings.Split(s, " ")
	for index, word := range words {
		if len(word) == 0 || word[0] != '$' {
			continue
		}
		if word[0] == '$' {
			words[index] = f(word[1:])
		}
	}

	return strings.Join(words, " ")
}

func addSan(s string) string {
	return fmt.Sprintf("%s-san", s)
}
