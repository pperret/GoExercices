// intmap manipulates a set of integers as a map
package main

import (
	"bytes"
	"fmt"
	"sort"
)

// An IntMap is a set of integers stored as a map.
type IntMap map[int]bool

// Has reports whether the map contains the value x.
func (m *IntMap) Has(x int) bool {
	return (*m)[x]
}

// Add adds the value x to the map.
func (m *IntMap) Add(x int) {
	(*m)[x] = true
}

// UnionWith sets m to the union of m and n.
func (m *IntMap) UnionWith(n *IntMap) {
	for v := range *n {
		(*m)[v] = true
	}
}

// String returns the map as a string of the form "{1 2 3}".
// As the order in map is random: a conversion in a sorted slice is performed beforehand
func (m *IntMap) String() string {

	// Create the temporary slice
	var t []int
	for v := range *m {
		t = append(t, v)
	}

	// Sort the slice
	sort.Ints(t)

	// Generate the string
	var buf bytes.Buffer
	buf.WriteByte('{')
	for _, v := range t {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte('}')
	return buf.String()
}
