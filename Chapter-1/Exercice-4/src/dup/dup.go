// Print the names of all files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Number of line occurrences
	counts := make(map[string]int)
	// File names (slice of file names) for each line
	names := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "-", counts, names)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts, names)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t", n, line)
			sep := ""
			for _, name := range names[line] {
				fmt.Printf("%s%s", sep, name)
				sep = " "
			}
			fmt.Println()
		}
	}
}

// Analyse one file
func countLines(f *os.File, name string, counts map[string]int, names map[string][]string) {
	already := make(map[string]bool)
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if already[line] == false {
			already[line] = true
			names[line] = append(names[line], name)
		}
	}
}
