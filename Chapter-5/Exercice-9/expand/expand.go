// expand is a package to replace variables in strings
package main

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// expand replaces variables in a string using the provided function to convert variable into value
func expand(s string, f func(string) string) string {

	// Look for a variable name (beginning with the $ character)
	// The returned index is in bytes
	i := strings.IndexRune(s, '$')
	if i == -1 {
		return s
	}

	// Save the index (in bytes) of the variable in the string
	varIndex := i

	// Get the variable name
	name := ""
	for i++; i < len(s); {
		c, l := utf8.DecodeRuneInString(s[i:])
		if !unicode.IsLetter(c) {
			break
		}
		i += l
		name += string(c)
	}

	// Call the function recusively
	return s[:varIndex] + f(name) + expand(s[i:], f)
}
