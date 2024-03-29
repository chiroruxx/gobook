package toposort

import (
	"sort"
)

var PreReqs = map[string][]string{
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

func TopoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitALl func(items []string)

	visitALl = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitALl(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitALl(keys)
	return order
}
