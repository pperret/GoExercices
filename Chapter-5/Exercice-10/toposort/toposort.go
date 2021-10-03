// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
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
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

// topoSort sorts the courses
func topoSort(m map[string][]string) []string {
	var result []string
	seen := make(map[string]bool)

	var visitAll func(course string)
	visitAll = func(course string) {
		if !seen[course] {
			for _, c := range m[course] {
				visitAll(c)
			}
			seen[course] = true
			result = append(result, course)
		}
	}

	for course := range m {
		visitAll(course)
	}
	return result
}
