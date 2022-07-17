// intset manipulates a set of integers as a bit array
package main

import (
	"bytes"
	"fmt"
)

// An IntSet32 is a set of small nonnegative integers.
// Its zero value represents the empty set.
type IntSet32 struct {
	words []uint32
}

// uint32Size is the size of a uint32 in bits
const uint32Size = 32 << (^uint64(0) >> 63)

// Dup duplicates a set of integers
func (s *IntSet32) Dup() *IntSet32 {
	var r IntSet32
	r.words = append(r.words, s.words...)
	return &r
}

// Has reports whether the set contains the nonnegative value x.
func (s *IntSet32) Has(x int) bool {
	word, bit := x/uint32Size, uint(x%uint32Size)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the nonnegative value x to the set.
func (s *IntSet32) Add(x int) {
	word, bit := x/uint32Size, uint(x%uint32Size)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet32) UnionWith(t *IntSet32) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet32) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uint32Size; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", uint32Size*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
