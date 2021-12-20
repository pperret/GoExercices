// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

// echo returns third times the incoming string
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

// handleConn manages an incoming connection
func handleConn(c net.Conn) {
	var wg sync.WaitGroup // number of working goroutines
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			echo(c, input.Text(), 1*time.Second)
		}()
	}
	// NOTE: ignoring potential errors from input.Err()

	// Wait for completion
	wg.Wait()

	// Close the connection
	c.Close()
}

// main is the entry point of the program
func main() {

	// Listen for incoming connections
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	// Accept incoming connections
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
