package practice5_11

import (
	"fmt"
	"sort"
)

func TopoSort(m map[string][]string) ([]string, error) {
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
