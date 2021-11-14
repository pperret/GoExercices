// Xmlselect prints the text of selected elements of an XML document.
// exemple: curl -s http://www.w3.org/TR/2006/RECxml1120060816 | $GOPATH/bin/xmlselect div div h2
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

// main is the entry point of the program
func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string          // stack of element names
	var attrs []map[string]bool // stack of element names and their attributes
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push name
			item := make(map[string]bool)
			item[tok.Name.Local] = true
			for _, attr := range tok.Attr {
				item[attr.Name.Local] = true
			}
			attrs = append(attrs, item) // push attributes
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop name
			attrs = attrs[:len(attrs)-1] // pop attributes
		case xml.CharData:
			if containsAll(attrs, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []map[string]bool, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0][y[0]] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
