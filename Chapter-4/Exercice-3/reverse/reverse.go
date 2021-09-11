// Reverse an array of integer
package main

import (
	"fmt"
)

// main is the entry point of the program
func main() {
	tab := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	reverse(&tab)

	fmt.Println(tab)
}

// reverse reverses a array of ints in place.
func reverse(s *[9]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
