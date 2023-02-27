package main

import (
	"bytes"
	"fmt"
)

type IntSet2 struct {
	words []uint64
}

func (s *IntSet2) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet2) Len() int {
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

func (s *IntSet2) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet2) Remove(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] &^= 1 << bit
}

func (s *IntSet2) Clear() {
	s.words = []uint64{}
}

func (s *IntSet2) Copy() *IntSet2 {
	var copied IntSet2

	for _, word := range s.words {
		copied.words = append(copied.words, word)
	}

	return &copied
}

func (s *IntSet2) UnionWith(t *IntSet2) {
	for i, tWord := range t.words {
		if i < len(s.words) {
			s.words[i] |= tWord
		} else {
			s.words = append(s.words, tWord)
		}
	}
}

func (s *IntSet2) String() string {
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
	var x IntSet2
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y := x.Copy()
	x.Clear()

	fmt.Println(x.String(), y.String())
}
