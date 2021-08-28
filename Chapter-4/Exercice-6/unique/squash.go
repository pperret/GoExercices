// Remove duplicate adjacent spaces in a byte string (UTF-8 encoded)
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	str1 := "	 Bonjour    à tous les éléments  !!      !! 	 "
	fmt.Printf("%q\n", str1)

	tab1 := []byte(str1)
	fmt.Printf("Before: % x\n", tab1)

	tab2 := squash(tab1)
	fmt.Printf("After: % x\n", tab2)

	str2 := string(tab2)
	fmt.Printf("%q\n", str2)
}

// Squashes adjacent spaces
func squash(s []byte) []byte {
	to := 0
	space := false
	for from := 0; from < len(s); {
		r, l := utf8.DecodeRune(s[from:])
		if unicode.IsSpace(r) {
			if !space {
				s[to] = ' '
				to++
			}
			space = true
		} else {
			copy(s[to:], s[from:from+l])
			to += l
			space = false
		}
		from += l
	}
	return s[:to]
}
