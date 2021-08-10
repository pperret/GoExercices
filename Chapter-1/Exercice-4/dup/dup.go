// Prints the names of files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Count of each line
	counts := make(map[string]int)
	// File names (slice of file names) for each line
	ins := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "-", counts, ins)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts, ins)
			f.Close()
		}
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\t", count, line)
			sep := ""
			for _, name := range ins[line] {
				fmt.Printf("%s%s", sep, name)
				sep = ", "
			}
			fmt.Println()
		}
	}
}

// Analyzes one file
func countLines(f *os.File, name string, counts map[string]int, ins map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !searchAlreadyFound(name, ins[line]) {
			ins[line] = append(ins[line], name)
		}
	}
}

// Searchs if file name is already in the list of file names
func searchAlreadyFound(name string, names []string) bool {
	for _, n := range names {
		if name == n {
			return true
		}
	}
	return false
}
