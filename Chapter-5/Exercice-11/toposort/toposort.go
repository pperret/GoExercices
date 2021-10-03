// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"os"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},

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
	courses, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	for i, course := range courses {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

// courseState is the analyse state of a course
type courseState int

const (
	courseNotSeen courseState = iota
	coursePending
	courseSeen
)

// topoSort sorts the courses
func topoSort(m map[string][]string) ([]string, error) {
	var result []string
	seen := make(map[string]courseState)

	var visitAll func(course string) error
	visitAll = func(course string) error {
		switch seen[course] {
		case courseNotSeen:
			seen[course] = coursePending
			for _, c := range m[course] {
				err := visitAll(c)
				if err != nil {
					return err // Cycle
				}
			}
			seen[course] = courseSeen
			result = append(result, course)
			return nil // Done
		case coursePending:
			return fmt.Errorf("cycle detected") // Cycle
		case courseSeen:
			return nil // Already seen
		}
		return fmt.Errorf("invalid course state")
	}

	for course := range m {
		err := visitAll(course)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
