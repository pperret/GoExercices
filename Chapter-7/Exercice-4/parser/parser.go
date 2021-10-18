// parser prints the links in an HTML document read from a string reader
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/net/html"
)

// StringReader is my own StringReader implementation
type StringReader struct {
	str string
	pos int
}

// Read is the implementation of the io.Reader interface for StringReader
func (msr *StringReader) Read(p []byte) (lg int, err error) {
	lg = copy(p, msr.str[msr.pos:])
	msr.pos += lg
	if msr.pos >= len(msr.str) {
		err = io.EOF
	}
	return
}

// NewStringReader creates a custom implementation of strings.NewReader
func NewStringReader(str string) io.Reader {
	return &StringReader{str: str}
}

// main is the entry point of the program
func main() {
	// Read the HTML document from stdin
	text, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read text %v\n", err)
		os.Exit(1)
	}

	// Create the string reader
	reader := NewStringReader(string(text))

	// Parse the HTML docuent
	doc, err := html.Parse(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse HTML %v\n", err)
		os.Exit(1)
	}

	// Display the result
	html.Render(os.Stdout, doc)
}
