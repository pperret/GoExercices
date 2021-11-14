// Xmltree prints the tree view of a XML document.
// exemple: curl -s http://www.w3.org/TR/2006/RECxml1120060816 | $GOPATH/bin/xmltree
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Node interface{} // CharData or *Element
type CharData string
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

// main is the entry point of the program
func main() {
	root, err := parseXmlTree(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse document: %v\n", err)
		os.Exit(1)
	}
	printXmlTree(root, 0)
}

// parseXmlTree parse a XML document and generates the corresponding tree
func parseXmlTree(reader io.Reader) (Node, error) {
	dec := xml.NewDecoder(os.Stdin)
	var stack []*Element
	var root Node

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			element := &Element{Type: tok.Name, Attr: tok.Attr}
			if len(stack) == 0 {
				root = element
			} else {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, element)
			}
			stack = append(stack, element) // push the element
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop element
		case xml.CharData:
			str := strings.TrimSpace(string(tok))
			if len(str) == 0 {
				continue
			}
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, CharData(str))
			}
		}
	}

	return root, nil
}

// printXmlTree displays the XML tree
func printXmlTree(node Node, level int) {
	switch n := node.(type) {
	case *Element:
		fmt.Printf("%*s<%s", level*2, "", n.Type.Local)
		for _, attr := range n.Attr {
			fmt.Printf(" %s=%q", attr.Name.Local, attr.Value)
		}
		if len(n.Children) == 0 {
			fmt.Print("/>\n")

		} else {
			fmt.Print(">\n")
			for _, child := range n.Children {
				printXmlTree(child, level+1)
			}
			fmt.Printf("%*s</%s>\n", level*2, "", n.Type.Local)
		}
	case CharData:
		fmt.Printf("%*s%s\n", level*2, "", n)
	default:
		log.Fatalf("Invalid type: %T", node)
	}
}
