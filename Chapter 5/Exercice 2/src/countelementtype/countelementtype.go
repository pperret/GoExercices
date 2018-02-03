// CountElementType prints the number of each element type in a HTML document
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// main is the entry point of the program
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse HML document: %v\n", err)
		os.Exit(1)
	}

	// Create a map for the number of each node type
	m := make(map[string]int)

	// Analyze the HTML document
	visit(m, doc)

	// Display the result
	for t, n := range m {
		fmt.Printf("%s: %d\n", t, n)
	}
}

// visit counts the number of each element type
func visit(m map[string]int, node *html.Node) {
	if node.Type == html.ElementNode {
		m[node.Data]++
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		visit(m, child)
	}
}
