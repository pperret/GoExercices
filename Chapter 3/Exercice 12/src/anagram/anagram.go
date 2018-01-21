// Check if two strings are anagrams
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <word> <word>\n", os.Args[0])
		os.Exit(1)
	}

	res := anagram(os.Args[1], os.Args[2])
	if res == true {
		fmt.Println("There are anagrams")
	} else {
		fmt.Println("There are not anagrams")
	}
}

func anagram(s1, s2 string) bool {
	so1 := order(s1)
	so2 := order(s2)

	return so1 == so2
}

func order(s string) string {
	t := []rune(s)
	for j := 0; j < len(t)-1; j++ {
		for i := j + 1; i < len(t); i++ {
			if t[j] > t[i] {
				t[j], t[i] = t[i], t[j]
			}
		}
	}
	return string(t)
}
