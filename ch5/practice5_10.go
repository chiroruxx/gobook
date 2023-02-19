package main

import (
	"fmt"
)

var preReqs2 = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort2(preReqs2) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort2(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	var visit func(item string)
	visit = func(item string) {
		if !seen[item] {
			seen[item] = true
			for _, dependent := range m[item] {
				visit(dependent)
			}
			order = append(order, item)
		}
	}

	visitALl := func(m map[string][]string) {
		for item := range m {
			visit(item)
		}
	}

	visitALl(m)
	return order
}
