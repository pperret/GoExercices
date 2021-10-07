package main

import (
	"fmt"
)

// min computes the minimum value of a list of integers
// An error is returned if the list is empty
func min(vals ...int) (int, bool) {
	if len(vals) == 0 {
		return 0, false
	}
	m := vals[0]
	for _, v := range vals[1:] {
		if v < m {
			m = v
		}
	}
	return m, true
}

// min1 computes the minimum value of a list of integers
// One value must be provided
func min1(m int, vals ...int) int {
	for _, v := range vals {
		if v < m {
			m = v
		}
	}
	return m
}

// max computes the maximum value of a list of integers
// An error is returned if the list is empty
func max(vals ...int) (int, bool) {
	if len(vals) == 0 {
		return 0, false
	}
	m := vals[0]
	for _, v := range vals[1:] {
		if v > m {
			m = v
		}
	}
	return m, true
}

// max1 computes the maximum value of a list of integers
// One value must be provided
func max1(m int, vals ...int) int {
	for _, v := range vals {
		if v > m {
			m = v
		}
	}
	return m
}

// main is the entry point of the program
func main() {
	if m, ok := min(0, 3, 5, -2); ok {
		fmt.Printf("Min: %d\n", m)
	}

	if m, ok := max(0, 3, 5, -2); ok {
		fmt.Printf("Max: %d\n", m)
	}

	fmt.Printf("Min1: %d\n", min1(0, 3, 5, -2))

	fmt.Printf("Max1: %d\n", max1(0, 3, 5, -2))
}
