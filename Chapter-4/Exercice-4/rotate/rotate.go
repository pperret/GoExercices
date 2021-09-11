// Rotate a slice of integers
package main

import (
	"fmt"
)

// main is the entry point of the program
func main() {
	tab := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	tab = rotate(tab, 2)

	fmt.Println(tab)
}

// rotate rotates a slice of ints in place
func rotate(s []int, p int) []int {
	t := make([]int, p)
	copy(t, s[0:p])
	copy(s, s[p:])
	copy(s[len(s)-p:], t)
	return s
}
