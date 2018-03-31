// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
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
	courses := topoSort(prereqs)
	if courses == nil {
		fmt.Printf("Cycle detected!!\n")
	} else {
		for i, course := range courses {
			fmt.Printf("%d:\t%s\n", i+1, course)
		}
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
func topoSort(m map[string][]string) []string {
	var result []string
	seen := make(map[string]courseState)
	var visitAll func(course string) bool

	visitAll = func(course string) bool {
		switch seen[course] {
		case courseNotSeen:
			seen[course] = coursePending
			for _, c := range m[course] {
				f := visitAll(c)
				if f == false {
					return false // Cycle
				}
			}
			seen[course] = courseSeen
			result = append(result, course)
			return true // Done
		case coursePending:
			return false // Cycle
		case courseSeen:
			return true // Already seen
		}
		return false
	}

	for course := range m {
		f := visitAll(course)
		if f == false {
			return nil
		}
	}
	return result
}
