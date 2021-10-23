// limitreader is a sample program implementing the LimitReader
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// limitReader is a standard reader limiting the number of available bytes
type limitReader struct {
	reader io.Reader
	len    int64
	pos    int64
}

// Read is the implementation of the io.Reader interface for limitReader
func (lr *limitReader) Read(p []byte) (lg int, err error) {

	// Compute the number of bytes to read
	max := int64(len(p))
	if lr.pos+int64(max) > lr.len {
		max = lr.len - lr.pos
	}

	// Read the bytes
	lg, err = lr.reader.Read(p[:max])

	// Compute the error code and update the position
	if err == nil || err == io.EOF {
		lr.pos += int64(lg)
		if lr.pos >= lr.len {
			err = io.EOF
		}
	}

	return
}

// LimitReader creates an instance of limitReader
func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{reader: r, len: n, pos: 0}
}

// main is the entry point of the program
func main() {

	// Create the LimitReader
	reader := LimitReader(os.Stdin, 1000)

	// Read the content
	bytes, err := ioutil.ReadAll(reader)

	// Display the result
	fmt.Printf("len=%d, err=%v\n", len(bytes), err)
}
