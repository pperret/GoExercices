// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// client data
type ClientChannel chan<- string // outgoing message channel
type ClientData struct {
	outgoing ClientChannel // outgoing message channel
	id       string        // Client's name
}

var (
	entering = make(chan ClientData)
	leaving  = make(chan ClientChannel)
	messages = make(chan string) // all incoming client messages
)

// broadcaster sends messages to each connected client
func broadcaster() {
	clients := make(map[ClientChannel]string) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for client := range clients {
				client <- msg
			}

		case clientData := <-entering:
			// Message is sent before adding the new client to the list
			// So, he/she is not in the online list
			for _, name := range clients {
				clientData.outgoing <- name + " is online"
			}
			clients[clientData.outgoing] = clientData.id

		case clientChannel := <-leaving:
			delete(clients, clientChannel)
			close(clientChannel)
		}
	}
}

//!-broadcaster

// handleConn manages a connection with a user
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ClientData{ch, who}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

// clientWriter is a go routine to send message to a user
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

// main is the entry point of the program
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
