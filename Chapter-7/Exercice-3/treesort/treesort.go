// tree is a sample of tree of integers
package main

import (
	"fmt"
)

// tree is a node of a tree of integers
type tree struct {
	value       int
	left, right *tree
}

// main is the entry point of the program
func main() {
	values := []int{10, 5, 100, 800, 700, 600, 2000, -3, -10000}
	Sort(values)
	fmt.Println(values)
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}

	// Print the internal structure of the tree
	fmt.Println(root)
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// add inserts an integer at the right place into the tree
func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// String converts the tree in a string revealing its internal structure
func (t *tree) String() string {
	return t.stringNode(0, "")
}

// stringNode stringyfies a subtree.
func (t *tree) stringNode(level int, sep string) string {
	if t == nil {
		return ""
	}
	s := ""
	if t.left != nil {
		s += t.left.stringNode(level+1, "\u250c") + "\n"
	}
	s += fmt.Sprintf("% *s%d", level*2, sep, t.value)
	if t.right != nil {
		s += "\n" + t.right.stringNode(level+1, "\u2514")
	}
	return s
}
