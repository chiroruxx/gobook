package main

import (
	"bytes"
	"fmt"
)

type IntSet5 struct {
	words []uint64
}

func (s *IntSet5) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet5) Len() int {
	var count int
	for word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			count += word & (1 << j)
		}
	}

	return count
}

func (s *IntSet5) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet5) AddAll(x ...int) {
	for _, xItem := range x {
		s.Add(xItem)
	}
}

func (s *IntSet5) Remove(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] &^= 1 << bit
}

func (s *IntSet5) Clear() {
	s.words = []uint64{}
}

func (s *IntSet5) Copy() *IntSet5 {
	var copied IntSet5

	for _, word := range s.words {
		copied.words = append(copied.words, word)
	}

	return &copied
}

func (s *IntSet5) UnionWith(t *IntSet5) {
	for i, tWord := range t.words {
		if i < len(s.words) {
			s.words[i] |= tWord
		} else {
			s.words = append(s.words, tWord)
		}
	}
}

func (s *IntSet5) IntersectWith(t *IntSet5) {
	if len(s.words) > len(t.words) {
		s.words = s.words[:len(t.words)]
	}

	for i, tWord := range t.words {
		if i < len(s.words) {
			s.words[i] &= tWord
		}
	}
}

func (s *IntSet5) DifferenceWith(t *IntSet5) {
	for i, tWord := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tWord
		}
	}
}

func (s *IntSet5) SymmetricDifference(t *IntSet5) {
	for i, tWord := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tWord
		} else {
			s.words = append(s.words, tWord)
		}
	}
}

func (s *IntSet5) Elems() []int {
	var elems []int

	for i, word := range s.words {
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

func (s *IntSet5) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for _, word := range s.Elems() {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", word)
	}

	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x IntSet5
	x.AddAll(1, 144, 9, 12, 145)
	fmt.Println(x.String())
}
