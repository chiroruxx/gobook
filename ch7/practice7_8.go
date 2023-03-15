package main

import (
	"fmt"
	"gobook/ch7/practice7_8"
	"sort"
)

func main() {
	items := []*practice7_8.Item{
		{1, "pencil", "Mark"},
		{2, "pencil", "John"},
		{3, "joint", "John"},
		{4, "joint", "Mark"},
	}
	history := make([]practice7_8.SortKey, 0)
	practice7_8.PrintItems(items)

	fmt.Println()
	practice7_8.ClickName(&history)
	sort.Sort(practice7_8.ByLastClick{Items: items, History: history})
	practice7_8.PrintItems(items)

	fmt.Println()
	practice7_8.ClickOwner(&history)
	sort.Sort(practice7_8.ByLastClick{Items: items, History: history})
	practice7_8.PrintItems(items)
}
