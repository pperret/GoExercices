// palindrome is a sample program using the sort interface to check if a string is a palindrome
package main

import (
	"fmt"
	"os"
	"sort"
)

// IsPalindrome check if the provided list is a palindrome
func IsPalindrome(list sort.Interface) bool {
	for i, j := 0, list.Len()-1; i < j; i, j = i+1, j-1 {
		if list.Less(i, j) || list.Less(j, i) {
			return false
		}
	}
	return true
}

// runeSlice is slice of runes implementing the sort.Interface
type runeSlice []rune

func (rs runeSlice) Len() int           { return len(rs) }
func (rs runeSlice) Less(i, j int) bool { return rs[i] < rs[j] }
func (rs runeSlice) Swap(i, j int)      { rs[i], rs[j] = rs[j], rs[i] }

// main is the entry point of the program
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <string> [strings...]\n", os.Args[0])
		os.Exit(1)
	}
	// Check if each argument is a palindrome
	for _, str := range os.Args[1:] {
		var rs runeSlice = []rune(str)
		switch IsPalindrome(rs) {
		case true:
			fmt.Printf("%q is a palindrome\n", str)
		case false:
			fmt.Printf("%q is not a palindrome\n", str)
		}
	}
}
