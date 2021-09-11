// Remove dupplicate adjacent strings in a slice
package main

import (
	"fmt"
)

// main is the entry point of the program
func main() {
	tab := []string{"toto", "titi", "tutu", "tutu", "tete", "tete", "tete", "tata", "titi"}

	tab2 := adjacent(tab)

	fmt.Println(tab)
	fmt.Println(tab2)
}

// adjacent purges adjacent strings
func adjacent(s []string) []string {
	i := 0
	for _, str := range s {
		if i == 0 || str != s[i-1] {
			s[i] = str
			i++
		}
	}
	return s[:i]
}
