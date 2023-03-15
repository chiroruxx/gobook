package practice6_2

import (
	"gobook/ch6/practice6_1"
)

type IntSet struct {
	practice6_1.IntSet
}

func (s *IntSet) Copy() *IntSet {
	result := IntSet{*s.IntSet.Copy()}
	return &result
}

func (s *IntSet) UnionWith(t *IntSet) {
	s.IntSet.UnionWith(&t.IntSet)
}

func (s *IntSet) AddAll(x ...int) {
	for _, xItem := range x {
		s.Add(xItem)
	}
}
