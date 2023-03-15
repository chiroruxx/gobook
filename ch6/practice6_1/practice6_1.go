package practice6_1

import "gobook/ch6/intset"

type IntSet struct {
	intset.IntSet
}

func (s *IntSet) Len() int {
	var count int
	for word := range s.Words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			count += word & (1 << j)
		}
	}

	return count
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.Words) {
		s.Words = append(s.Words, 0)
	}
	s.Words[word] &^= 1 << bit
}

func (s *IntSet) Clear() {
	s.Words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	var copied IntSet

	for _, word := range s.Words {
		copied.Words = append(copied.Words, word)
	}

	return &copied
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tWord := range t.Words {
		if i < len(s.Words) {
			s.Words[i] |= tWord
		} else {
			s.Words = append(s.Words, tWord)
		}
	}
}
