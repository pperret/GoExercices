// expand is a package to replace variables in strings
package main

import (
	"fmt"
	"testing"
)

// f1 is the remplacement function
func f1(s string) string {
	if s == "foo" {
		return "1"
	} else if s == "var" {
		return "2"
	} else {
		return "Unknown"
	}
}

// entry is an elementary test value
type entry struct {
	// value is the source string
	value string
	// expected is the expected result after applying expand
	expected string
}

// List of test values
var tests = []entry {
	{"$foo", "1"},
	{"$var", "2"},
	{"$foo$var", "12"},
	{"$var$foo", "21"},
	{"$dummy", "Unknown"},
}

// TestExpand1 tests the expand function
func TestExpand1(t *testing.T) {

	// Loop on elementary tests
	for _, e := range tests {
		s := expand(e.value, f1)
		if s != e.expected {
			fmt.Printf("Error: source=%s, expected=%s, found=%s\n", e.value, e.expected, s)
		}
	}
}
