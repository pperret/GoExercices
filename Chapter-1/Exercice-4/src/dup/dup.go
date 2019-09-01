// Print the names of all files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// File names (slice of file names) for each line
	counts := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "-", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}
	for line, names := range counts {
		if len(names) > 1 {
			fmt.Printf("%d\t%s\t", len(names), line)
			sep := ""
			for _, name := range names {
				fmt.Printf("%s%s", sep, name)
				sep = " "
			}
			fmt.Println()
		}
	}
}

// Analyse one file
func countLines(f *os.File, name string, counts map[string][]string) {
	already := make(map[string]bool)
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if already[line] == false {
			already[line] = true
			counts[line] = append(counts[line], name)
		}
	}
}
