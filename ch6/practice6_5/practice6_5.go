package practice6_5

import (
	"gobook/ch6/practice6_4"
)

const bitSize = 32 << (^uint(0) >> 63)

type IntSet struct {
	practice6_4.IntSet
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/bitSize, uint(x%bitSize)
	return word < len(s.Words) && s.Words[word]&(1<<bit) != 0
}

func (s *IntSet) Len() int {
	var count int
	for word := range s.Words {
		if word == 0 {
			continue
		}

		for j := 0; j < bitSize; j++ {
			count += word & (1 << j)
		}
	}

	return count
}

func (s *IntSet) Add(x int) {
	word, bit := x/bitSize, uint(x%bitSize)
	for word >= len(s.Words) {
		s.Words = append(s.Words, 0)
	}
	s.Words[word] |= 1 << bit
}

func (s *IntSet) Remove(x int) {
	word, bit := x/bitSize, uint(x%bitSize)
	for word >= len(s.Words) {
		s.Words = append(s.Words, 0)
	}
	s.Words[word] &^= 1 << bit
}

func (s *IntSet) Clear() {
	s.Words = []uint64{}
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
		for j := 0; j < bitSize; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, i*bitSize+j)
			}
		}
	}

	return elems
}
