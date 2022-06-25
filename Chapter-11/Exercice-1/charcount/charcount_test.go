package main

import (
	"strings"
	"testing"
	"unicode/utf8"
)

// TestCharcount is the test function of Charcount function
func TestCharcount(t *testing.T) {
	// Tests list
	var tests = []struct {
		input   string               // String to be analyzed
		count   map[rune]int         // Count of each character
		utflen  [utf8.UTFMax + 1]int // Count by character size
		letter  int                  // Count of letters
		digit   int                  // Count of digits
		space   int                  // Count of spaces
		invalid int                  // Count of invalid character
	}{
		{
			"Try again...",
			map[rune]int{'T': 1, 'r': 1, 'y': 1, ' ': 1, 'a': 2, 'g': 1, 'i': 1, 'n': 1, '.': 3},
			[utf8.UTFMax + 1]int{0, 12, 0, 0, 0},
			8, 0, 1, 0},
		{
			"été",
			map[rune]int{'é': 2, 't': 1},
			[utf8.UTFMax + 1]int{0, 1, 2, 0, 0},
			3, 0, 0, 0},
		{
			"Hi, 世.",
			map[rune]int{'H': 1, 'i': 1, ',': 1, ' ': 1, '世': 1, '.': 1},
			[utf8.UTFMax + 1]int{0, 5, 0, 1, 0},
			3, 0, 1, 0},
	}

	// Loop on tests
	for _, test := range tests {

		// Create a reader containing the input string
		s := strings.NewReader(test.input)

		// Analyze the string
		count, utflen, letter, digit, space, invalid, err := Charcount(s)
		if err != nil {
			t.Errorf("Error: %v", err)
			continue
		}

		// Check the count of each character
		for r, c := range count {
			if test.count[r] != c {
				t.Errorf("Bad count in '%s' for character '%c': expected: %d, found: %d", test.input, r, test.count[r], c)
			}
		}
		for r, c := range test.count {
			if count[r] == 0 {
				t.Errorf("Bad count in '%s' for character '%c': expected: %d, found: %d", test.input, r, c, 0)
			}
		}

		// Check the count by character size
		for i, c := range test.utflen {
			if utflen[i] != c {
				t.Errorf("Bad count in '%s' for character length of %d: expected: %d, found: %d", test.input, i, c, utflen[i])
			}
		}

		// Check the count of letters
		if letter != test.letter {
			t.Errorf("Bad letter count in '%s': expected: %d, found: %d", test.input, test.letter, letter)
		}

		// Check the count of digits
		if digit != test.digit {
			t.Errorf("Bad digit count in '%s': expected: %d, found: %d", test.input, test.digit, digit)
		}

		// Check the count of spaces
		if space != test.space {
			t.Errorf("Bad space count in '%s': expected: %d, found: %d", test.input, test.space, space)
		}

		// Check the count of invalid characters
		if invalid != test.invalid {
			t.Errorf("Bad invalid count in '%s': expected: %d, found: %d", test.input, test.invalid, invalid)
		}
	}
}
