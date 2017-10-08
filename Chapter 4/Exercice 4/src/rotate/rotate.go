// Rotate an slice of integer
package main

import (
	"fmt"
)

func main() {
	tab := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	tab = rotate(tab, 2)

	fmt.Println(tab)
}

// Rotate a slice of ints in place
func rotate(s []int, p int) []int {
	t := make([]int, len(s))
	for i := range s {
		t[p] = i
		p = (p + 1) % len(s)
	}
	return t
}
