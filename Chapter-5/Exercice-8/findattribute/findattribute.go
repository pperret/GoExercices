// findattribute looks for an element containing a specific attribute in a HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// main is the entry point of the program
func main() {
	// Check the argument count
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <url> <attribute>\n", os.Args[0])
		os.Exit(1)
	}

	// Retrieve the HTML document from the URL
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to access the URL\n")
		os.Exit(2)
	}
	defer resp.Body.Close()

	// Parse the HTML document
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse the document\n")
		os.Exit(3)
	}

	// Look for the attribute
	node := ElementByID(doc, os.Args[2])
	if node == nil {
		fmt.Fprintf(os.Stderr, "Attribute not found\n")
		os.Exit(4)
	}

	// Display the element containing the attribute
	fmt.Printf("Element name=%s\n", node.Data)
}

// ElementByID looks for an element attribute in the HTML document
func ElementByID(doc *html.Node, id string) *html.Node {

	// Enumerate HTML elements
	n := forEachNode(doc, startElement, nil, id)
	if n == nil {
		return nil
	}
	return n
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children is visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node, id string) bool, id string) *html.Node {

	// Call the pre-processing function
	if pre != nil {
		if !pre(n, id) {
			return n
		}
	}

	// Call recurvively the function for each child
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		n := forEachNode(c, pre, post, id)
		if n != nil {
			return n
		}
	}

	// Call the post-processing function
	if post != nil {
		if !post(n, id) {
			return n
		}
	}

	return nil
}

// startElement is called when an element starts to be parsed
func startElement(n *html.Node, id string) bool {

	// Check that the node is an element
	if n.Type != html.ElementNode {
		return true
	}

	// Look for the attribute in the element attributes list
	for _, a := range n.Attr {
		if a.Key == id {
			return false
		}
	}

	return true
}
