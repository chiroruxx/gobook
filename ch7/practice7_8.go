package main

import (
	"fmt"
	"sort"
)

type Item struct {
	Id    uint
	Name  string
	Owner string
}

type SortKey int

func (key *SortKey) isSame(one *Item, another *Item) bool {
	switch *key {
	case SortKeyName:
		return one.Name == another.Name
	case SortKeyOwner:
		return one.Owner == another.Owner
	}

	panic(key)
}

func (key *SortKey) less(one *Item, another *Item) bool {
	switch *key {
	case SortKeyName:
		return one.Name < another.Name
	case SortKeyOwner:
		return one.Owner < another.Owner
	}

	panic(key)
}

const (
	SortKeyName SortKey = iota
	SortKeyOwner
)

type byLastClick struct {
	items   []*Item
	history []SortKey
}

func (s byLastClick) Len() int {
	return len(s.items)
}

func (s byLastClick) Less(i, j int) bool {
	for index := len(s.history) - 1; index >= 0; index-- {
		key := s.history[index]
		if !key.isSame(s.items[i], s.items[j]) {
			return key.less(s.items[i], s.items[j])
		}
	}

	return false
}

func (s byLastClick) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}

func printItems(items []*Item) {
	fmt.Printf("%v\t%v\t%v\n", "Id", "Name", "Owner")
	for _, item := range items {
		fmt.Printf("%v\t%v\t%v\n", item.Id, item.Name, item.Owner)
	}
}

func clickName(history *[]SortKey) {
	*history = append(*history, SortKeyName)
}

func clickOwner(history *[]SortKey) {
	*history = append(*history, SortKeyOwner)
}

func main() {
	items := []*Item{
		{1, "pencil", "Mark"},
		{2, "pencil", "John"},
		{3, "joint", "John"},
		{4, "joint", "Mark"},
	}
	history := make([]SortKey, 0)
	printItems(items)

	fmt.Println()
	clickName(&history)
	sort.Sort(byLastClick{items: items, history: history})
	printItems(items)

	fmt.Println()
	clickOwner(&history)
	sort.Sort(byLastClick{items: items, history: history})
	printItems(items)
}
