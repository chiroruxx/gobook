package practice6_4

import (
	"gobook/ch6/practice6_3"
)

type IntSet struct {
	practice6_3.IntSet
}

func (s *IntSet) Copy() *IntSet {
	result := IntSet{*s.IntSet.Copy()}
	return &result
}

func (s *IntSet) UnionWith(t *IntSet) {
	s.IntSet.UnionWith(&t.IntSet)
}

func (s *IntSet) IntersectWith(t *IntSet) {
	s.IntSet.IntersectWith(&t.IntSet)
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	s.IntSet.DifferenceWith(&t.IntSet)
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	s.IntSet.SymmetricDifference(&t.IntSet)
}

func (s *IntSet) Elems() []int {
	var elems []int

	for i, word := range s.Words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, i*64+j)
			}
		}
	}

	return elems
}
