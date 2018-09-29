// Reverse an array of integer
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 := "Bonjour à toutes les bêtes"
	tab1 := []byte(str1)

	tab2 := reverseString(tab1)
	str2 := string(tab2)

	fmt.Println(str1)
	fmt.Printf("% x\n", tab1)
	fmt.Printf("% x\n", tab2)
	fmt.Println(str2)
}

func reverseByte(b []byte) {
	size := len(b)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[size-1-i] = b[size-1-i], b[i]
	}
}

// reverse a UTF8 string
func reverseString(str []byte) []byte {

	for i := 0; i < len(str); {
		_, size := utf8.DecodeRune(str[i:])
		reverseByte(str[i : i+size])
		i += size
	}
	reverseByte(str)
	return str
}
