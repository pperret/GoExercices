// Wordfreq computes counts of words.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// main is the entry point of the program
func main() {
	seen := make(map[string]int) // a set of words
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		seen[word]++
	}
	for word, freq := range seen {
		fmt.Printf("%s\t%d\n", word, freq)
	}
}
