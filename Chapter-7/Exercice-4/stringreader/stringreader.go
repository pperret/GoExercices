// stringreader is a sample program using a custom implementation of the StringReader interface
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/net/html"
)

// stringReader is a custom implementation of the StringReader class
type stringReader struct {
	str string
	pos int
}

// Read is the implementation of the io.Reader interface for stringReader
func (sr *stringReader) Read(p []byte) (lg int, err error) {
	lg = copy(p, sr.str[sr.pos:])
	sr.pos += lg
	if sr.pos >= len(sr.str) {
		err = io.EOF
	}
	return
}

// StringReader creates an instance of stringReader
func StringReader(str string) io.Reader {
	return &stringReader{str: str}
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
	reader := StringReader(string(text))

	// Parse the HTML docuent
	doc, err := html.Parse(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse HTML %v\n", err)
		os.Exit(1)
	}

	// Display the result
	html.Render(os.Stdout, doc)
}
