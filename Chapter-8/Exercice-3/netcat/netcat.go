// netcat is a console tool to connect to a TCP server
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// main is the entry point of the program
func main() {

	// Check arguments count
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <endpoint>\n", os.Args[0])
		return
	}

	// Resolve target server address
	addr, err := net.ResolveTCPAddr("tcp", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the server
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	// Create the channel
	done := make(chan struct{})

	// Start the go routine to display data returned by the server
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()

	// Send input data to the server up to the input stream is closed
	mustCopy(conn, os.Stdin)

	// Close the connection
	conn.CloseWrite()

	// wait for background goroutine to finish
	<-done

	// Close the read part of the connection
	conn.CloseRead()
}

// mustCopy copies data between a Reader and a Writer
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
