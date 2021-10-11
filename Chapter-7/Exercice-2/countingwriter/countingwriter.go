// countingwriter is a sample demonstrating a wrapper of the io.writer interface
// that counts the number of written bytes.
package main

import (
	"fmt"
	"io"
	"os"
)

// LocalWriter is a wrapper of io.Writer counting the written bytes
type LocalWriter struct {
	WrittenLength int64
	InnerWriter   io.Writer
}

// Write calls the inner writer and updates the counter
func (lw *LocalWriter) Write(p []byte) (int, error) {
	written, err := lw.InnerWriter.Write(p)
	if err == nil {
		lw.WrittenLength += int64(written)
	}
	return written, err
}

// CoutingWriter returns a new Writer that wraps the original io.Writer,
// and a pointer to an int64 variable that at any moment contains
// the number of bytes written to the new Writer.
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var lw LocalWriter
	lw.InnerWriter = w
	return &lw, &lw.WrittenLength
}

// main is the entry point of the program
func main() {
	lw, l := CountingWriter(os.Stdout)

	fmt.Fprintf(lw, "Hello world!!\n")

	fmt.Printf("len=%d\n", *l)
}
