// findtag looks for elements containing specific tags in a HTML document tree.
package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// main is the entry point of the program
func main() {
	// Check the argument count
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <url> [tags...]\n", os.Args[0])
		os.Exit(1)
	}

	// Retrieve the HTML document from the URL (TLS certificate errors are ignored)
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{Transport: customTransport}
	resp, err := client.Get(os.Args[1])
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

	// Look for the tags
	nodes := ElementsByTagName(doc, os.Args[2:]...)
	if nodes == nil {
		fmt.Fprintf(os.Stderr, "Error while seaching tags\n")
		os.Exit(4)
	}

	// Display the elements containing tags
	for _, node := range nodes {
		fmt.Printf("Element=%v\n", getNodePath(node))
	}
}

// ElementsByTagName looks for elements containing specific tags in the HTML document
func ElementsByTagName(doc *html.Node, tags ...string) []*html.Node {
	// For efficiency reasons in searches, tags are stored in a map
	ts := make(map[string]bool)
	for _, tag := range tags {
		ts[tag] = true
	}

	// Result list
	nodes := make([]*html.Node, 0)

	// Processing function for the current node
	processNode := func(node *html.Node) {

		// Check that the node is an element
		if node.Type == html.ElementNode {

			// Check if the element name is one of the tags
			if ts[node.Data] {
				nodes = append(nodes, node)
			}
		}
	}

	// Analyze the document recursively
	forEachNode(doc, processNode, nil)

	return nodes
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children is visited (preorder) and
// post is called after (postorder).
func forEachNode(node *html.Node, pre, post func(n *html.Node)) {

	// Call the pre-processing function for the current node
	if pre != nil {
		pre(node)
	}

	// Call recurvively the function for each child
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		forEachNode(child, pre, post)
	}

	// Call the post-processing function for the current node
	if post != nil {
		post(node)
	}
}

// getNodePath returns the "full path" of a HTML node
func getNodePath(node *html.Node) string {
	if node.Parent == nil || node.Parent.Type != html.ElementNode {
		return node.Data
	} else {
		return getNodePath(node.Parent) + "." + node.Data
	}
}
