package practice6_3

import (
	"gobook/ch6/practice6_2"
)

type IntSet struct {
	practice6_2.IntSet
}

func (s *IntSet) Copy() *IntSet {
	result := IntSet{*s.IntSet.Copy()}
	return &result
}

func (s *IntSet) UnionWith(t *IntSet) {
	s.IntSet.UnionWith(&t.IntSet)
}

func (s *IntSet) IntersectWith(t *IntSet) {
	if len(s.Words) > len(t.Words) {
		s.Words = s.Words[:len(t.Words)]
	}

	for i, tWord := range t.Words {
		if i < len(s.Words) {
			s.Words[i] &= tWord
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tWord := range t.Words {
		if i < len(s.Words) {
			s.Words[i] &^= tWord
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tWord := range t.Words {
		if i < len(s.Words) {
			s.Words[i] ^= tWord
		} else {
			s.Words = append(s.Words, tWord)
		}
	}
}
