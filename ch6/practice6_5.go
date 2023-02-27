package main

import (
	"bytes"
	"fmt"
)

const bitSize = 32 << (^uint(0) >> 63)

type IntSet6 struct {
	words []uint
}

func (s *IntSet6) Has(x int) bool {
	word, bit := x/bitSize, uint(x%bitSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet6) Len() int {
	var count int
	for word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < bitSize; j++ {
			count += word & (1 << j)
		}
	}

	return count
}

func (s *IntSet6) Add(x int) {
	word, bit := x/bitSize, uint(x%bitSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet6) AddAll(x ...int) {
	for _, xItem := range x {
		s.Add(xItem)
	}
}

func (s *IntSet6) Remove(x int) {
	word, bit := x/bitSize, uint(x%bitSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] &^= 1 << bit
}

func (s *IntSet6) Clear() {
	s.words = []uint{}
}

func (s *IntSet6) Copy() *IntSet6 {
	var copied IntSet6

	for _, word := range s.words {
		copied.words = append(copied.words, word)
	}

	return &copied
}

func (s *IntSet6) UnionWith(t *IntSet6) {
	for i, tWord := range t.words {
		if i < len(s.words) {
			s.words[i] |= tWord
		} else {
			s.words = append(s.words, tWord)
		}
	}
}

func (s *IntSet6) IntersectWith(t *IntSet6) {
	if len(s.words) > len(t.words) {
		s.words = s.words[:len(t.words)]
	}

	for i, tWord := range t.words {
		if i < len(s.words) {
			s.words[i] &= tWord
		}
	}
}

func (s *IntSet6) DifferenceWith(t *IntSet6) {
	for i, tWord := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tWord
		}
	}
}

func (s *IntSet6) SymmetricDifference(t *IntSet6) {
	for i, tWord := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tWord
		} else {
			s.words = append(s.words, tWord)
		}
	}
}

func (s *IntSet6) Elems() []int {
	var elems []int

	for i, word := range s.words {
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

func (s *IntSet6) String() string {
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
	var x IntSet6
	x.AddAll(1, 144, 9, 12, 145)
	fmt.Println(x.String())
}
