// Reverse a UTF-8 string
package main

import (
	"fmt"
	"unicode/utf8"
)

// main is the entry point of the program
func main() {
	str1 := "Bonjour à toutes les bêtes. ます !!"
	fmt.Println(str1)

	tab1 := []byte(str1)
	fmt.Printf("Before: % x\n", tab1)

	tab2 := reverseString(tab1)
	fmt.Printf("After: % x\n", tab2)

	str2 := string(tab2)
	fmt.Println(str2)
}

// reverseString reverses a UTF8 string
func reverseString(str []byte) []byte {

	for i := 0; i < len(str); {
		_, size := utf8.DecodeRune(str[i:])
		reverseByte(str[i : i+size])
		i += size
	}
	reverseByte(str)
	return str
}

// reverseByte reverses a slice of bytes
func reverseByte(b []byte) {
	size := len(b)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[size-1-i] = b[size-1-i], b[i]
	}
}
