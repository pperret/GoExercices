// Remove duplicate adjacent spaces in a byte string (UTF-8 encoded)
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	str1 := "Bonjour    à tous les éléments  !!      !!"
	tab1 := []byte(str1)

	tab2 := unique(tab1)

	str2 := string(tab2)

	fmt.Println(str1)
	fmt.Println(tab1)
	fmt.Println(tab2)
	fmt.Println(str2)
}

// Purge adjacent spaces
func unique(s []byte) []byte {
	to := 0
	space := false
	for from := 0; from < len(s); {
		r, l := utf8.DecodeRune(s[from:])
		if unicode.IsSpace(r) {
			if space == false {
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
