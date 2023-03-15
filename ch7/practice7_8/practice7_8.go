package practice7_8

import (
	"fmt"
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

type ByLastClick struct {
	Items   []*Item
	History []SortKey
}

func (s ByLastClick) Len() int {
	return len(s.Items)
}

func (s ByLastClick) Less(i, j int) bool {
	for index := len(s.History) - 1; index >= 0; index-- {
		key := s.History[index]
		if !key.isSame(s.Items[i], s.Items[j]) {
			return key.less(s.Items[i], s.Items[j])
		}
	}

	return false
}

func (s ByLastClick) Swap(i, j int) {
	s.Items[i], s.Items[j] = s.Items[j], s.Items[i]
}

func PrintItems(items []*Item) {
	fmt.Printf("%v\t%v\t%v\n", "Id", "Name", "Owner")
	for _, item := range items {
		fmt.Printf("%v\t%v\t%v\n", item.Id, item.Name, item.Owner)
	}
}

func ClickName(history *[]SortKey) {
	*history = append(*history, SortKeyName)
}

func ClickOwner(history *[]SortKey) {
	*history = append(*history, SortKeyOwner)
}
