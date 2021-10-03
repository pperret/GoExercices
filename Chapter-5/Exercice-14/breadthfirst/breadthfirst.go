// breadthfirst program prints the name of all the courses
package main

import (
	"fmt"
	"os"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

// main is the entry point of the program
func main() {
	var courses []string
	for course := range prereqs {
		courses = append(courses, course)
	}
	// Print courses breadth-first
	err := breadthFirst(prerequisites(), courses)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) error {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
	return nil
}

// prerequisites prints the course name and returns the prerequisites of a course
func prerequisites() func(string) []string {
	num := 0
	return func(course string) []string {
		num++
		fmt.Printf("%d - %s\n", num, course)
		return prereqs[course]
	}
}
