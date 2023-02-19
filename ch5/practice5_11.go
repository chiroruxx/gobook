package main

import (
	"fmt"
	"os"
	"sort"
)

var preReqs3 = map[string][]string{
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
	"linear algebra":        {"calculus"},
}

func main() {
	sorted, err := topoSort3(preReqs3)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	for i, course := range sorted {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort3(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string) error
	var pathList []string

	inArray := func(key string, items []string) bool {
		for _, item := range items {
			if item == key {
				return true
			}
		}
		return false
	}

	visitAll = func(items []string) error {
		for _, item := range items {
			if inArray(item, pathList) {
				return fmt.Errorf("loop occured!!")
			}

			if !seen[item] {
				seen[item] = true
				pathList = append(pathList, item)
				if err := visitAll(m[item]); err != nil {
					return err
				}
				order = append(order, item)
			}
		}

		pathList = nil
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	if err := visitAll(keys); err != nil {
		return nil, err
	}
	return order, nil
}
