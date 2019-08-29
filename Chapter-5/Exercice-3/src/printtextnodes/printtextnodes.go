// PrintTextNodes prints the content of each text node in a HTML document
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// main is the entry point of the program
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse HML document: %v\n", err)
		os.Exit(1)
	}

	// Analyze the HTML document
	visit(doc)
}

// visit prints the content of text nodes
func visit(node *html.Node) {

	// Display text node content if it is not empty
	if node.Type == html.TextNode {
		t := strings.TrimSpace(node.Data)
		if len(t) != 0 {
			fmt.Printf("%s\n", node.Data)
		}
	}

	// Recurse only for none script or style elements
	if node.Type != html.ElementNode || (node.Data != "script" && node.Data != "style") {
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			visit(child)
		}
	}
}
