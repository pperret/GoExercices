// expand is a package to replace variables in strings
package main

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// getRuneAtIndex returns the rune at the specified index in a string
func getRuneAtIndex(s string, i int) rune {
	return ([]rune(s))[i]
}

// expand replaces variables in a string using the provided function to convert variable into value
func expand(s string, f func(string) string) string {

	// Look for a variable name (begin by the $ character)
	i := strings.IndexRune(s, '$')
	if i == -1 {
		return s
	}

	// Get the variable name
	name := ""
	for i++ ; i<utf8.RuneCountInString(s) ; i++ {
		c := getRuneAtIndex(s, i)
		if unicode.IsLetter(c) == false {
			break
		}
		name += string(c)
	}

	// Replace the variable in the string using the provided function
	// Only once variable is replaced to prevent replacement of $fooo when the first matched variable name is $foo
	s = strings.Replace(s, "$" + name, f(name), 1)

	// Call the function recusively
	return expand(s, f)
}