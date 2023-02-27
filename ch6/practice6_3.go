package main

import (
	"bytes"
	"fmt"
)

type IntSet4 struct {
	words []uint64
}

func (s *IntSet4) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet4) Len() int {
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

func (s *IntSet4) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet4) AddAll(x ...int) {
	for _, xItem := range x {
		s.Add(xItem)
	}
}

func (s *IntSet4) Remove(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] &^= 1 << bit
}

func (s *IntSet4) Clear() {
	s.words = []uint64{}
}

func (s *IntSet4) Copy() *IntSet4 {
	var copied IntSet4

	for _, word := range s.words {
		copied.words = append(copied.words, word)
	}

	return &copied
}

func (s *IntSet4) UnionWith(t *IntSet4) {
	for i, tWord := range t.words {
		if i < len(s.words) {
			s.words[i] |= tWord
		} else {
			s.words = append(s.words, tWord)
		}
	}
}

func (s *IntSet4) IntersectWith(t *IntSet4) {
	if len(s.words) > len(t.words) {
		s.words = s.words[:len(t.words)]
	}

	for i, tWord := range t.words {
		if i < len(s.words) {
			s.words[i] &= tWord
		}
	}
}

func (s *IntSet4) DifferenceWith(t *IntSet4) {
	for i, tWord := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tWord
		}
	}
}

func (s *IntSet4) SymmetricDifference(t *IntSet4) {
	for i, tWord := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tWord
		} else {
			s.words = append(s.words, tWord)
		}
	}
}

func (s *IntSet4) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x, y IntSet4
	x.AddAll(1, 144, 9, 12, 145)
	y.AddAll(1, 144, 3, 12, 550)
	x.DifferenceWith(&y)
	fmt.Println(x.String())
}
