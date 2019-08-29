// Findlinks prints the links in an HTML document read from standard input.
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
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in node and returns the result.
func visit(links []string, node *html.Node) []string {

	// Map containing the link's attribute name for each targeted element
	var nodes = map[string]string{
		"a":      "href",
		"link":   "href",
		"iframe": "src",
		"script": "src",
		"img":    "src"}

	// Look for a link in the current element
	if node.Type == html.ElementNode {
		attributeName, ok := nodes[node.Data]
		if ok == true {
			// Too bad, it's not a map. I need to scan each attribute
			for _, a := range node.Attr {
				if a.Key == attributeName {
					links = append(links, a.Val)
				}
			}
		}
	}

	// Visit the childs (recursively)
	if node.FirstChild != nil {
		links = visit(links, node.FirstChild)
	}

	// Visit the siblings (recursively)
	if node.NextSibling != nil {
		links = visit(links, node.NextSibling)
	}

	return links
}
