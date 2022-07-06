// split is a test package intended to validate the string.Split function
package split

import (
	"strings"
	"testing"
)

// TestSplit tests the string.Split function using a tests table
func TestSplit(t *testing.T) {
	// Tests table for the string.Split function
	var tests = []struct {
		str  string // String to split
		sep  string // Separator
		want int    // Expected count of sub-strings
	}{
		{"a:b:c", ":", 3},
		{"abc", ":", 1},
		{"abc:def", ":", 2},
	}

	// Loop on the tests
	for _, test := range tests {
		words := strings.Split(test.str, test.sep)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d",
				test.str, test.sep, got, test.want)
		}
		// ...
	}
}
