// join is a variadic version of strings.Join
package main

import (
	"fmt"
	"os"
	"strings"
)

// join computes the concatenation of strings using a specified separator
func join(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	var res strings.Builder
	res.WriteString(strs[0])
	for _, s := range strs[1:] {
		res.WriteString(sep)
		res.WriteString(s)
	}
	return res.String()
}

// main is the entry point of the program
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <separator> [strings]\n", os.Args[0])
		os.Exit(1)
	}
	sep := os.Args[1]
	strs := os.Args[2:]
	fmt.Printf("%s\n", join(sep, strs...))
}
