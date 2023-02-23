// Package bzip provides a writer that uses bzip2 compression (bzip.org).
package bzip

import (
	"io"
	"os/exec"
)

type writer struct {
	stdin   io.WriteCloser
	command *exec.Cmd
}

// NewWriter returns a writer for bzip2-compressed streams
// It creates the underlying bzip2 command
func NewWriter(out io.Writer) (io.WriteCloser, error) {
	// Create the command
	cmd := exec.Command("/usr/bin/bzip2")

	// Create the input pipe
	in, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	// Assign the output stream
	cmd.Stdout = out

	// Start the command
	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	// Initialize and return the new writer
	return &writer{command: cmd, stdin: in}, nil
}

// Write sends data to the input pipe of the bzip2 command
func (w *writer) Write(data []byte) (int, error) {
	return w.stdin.Write(data)
}

// Close closes the input pipe and waits for commmand completion
func (w *writer) Close() error {
	// Close the input pipe
	err := w.stdin.Close()
	if err != nil {
		return err
	}

	// Wait for command completion
	return w.command.Wait()
}
