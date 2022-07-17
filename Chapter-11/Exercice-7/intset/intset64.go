// intset manipulates a set of integers as a bit array
package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small nonnegative integers.
// Its zero value represents the empty set.
type IntSet64 struct {
	words []uint64
}

// uint64Size is the size of a uint64 in bits
const uint64Size = 32 << (^uint64(0) >> 63)

// Dup duplicates a set of integers
func (s *IntSet64) Dup() *IntSet64 {
	var r IntSet64
	r.words = append(r.words, s.words...)
	return &r
}

// Has reports whether the set contains the nonnegative value x.
func (s *IntSet64) Has(x int) bool {
	word, bit := x/uint64Size, uint(x%uint64Size)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the nonnegative value x to the set.
func (s *IntSet64) Add(x int) {
	word, bit := x/uint64Size, uint(x%uint64Size)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet64) UnionWith(t *IntSet64) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet64) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uint64Size; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", uint64Size*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
