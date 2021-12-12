// ftpd is a minimal FTP server (RFC 959) supporting both IP V4/V6 addresses in active/passive mode
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// FtpConnection is the context of a FTP connection
type FtpConnection struct {
	tcpConnection      net.Conn
	user               string
	authenticated      bool
	writer             *bufio.Writer
	binary             bool
	passive            bool
	remoteDataEndPoint net.TCPAddr  // Used in active mode
	listener           net.Listener // Used in passive mode
	workingDirectory   string
}

// initialize initializes the connection context
func (conn *FtpConnection) initialize() error {
	remoteAddr := conn.tcpConnection.RemoteAddr()
	log.Printf("Incoming connection from %s", remoteAddr)

	// Parse the remote address
	tcpAddr, err := net.ResolveTCPAddr("tcp", remoteAddr.String())
	if err != nil {
		log.Printf("init: bad remote address: %s", remoteAddr)
		return fmt.Errorf("bad remote address: %s", remoteAddr)
	}

	// Initialize the default data endpoint
	conn.remoteDataEndPoint = *tcpAddr

	// Compute the default port
	conn.remoteDataEndPoint.Port = conn.remoteDataEndPoint.Port - 1

	conn.workingDirectory = "/tmp"
	conn.user = "anonymous"
	conn.writer = bufio.NewWriter(conn.tcpConnection)
	return nil
}

// writeString sends a reply to the client
func (conn *FtpConnection) writeString(format string, args ...interface{}) {
	conn.writer.WriteString(fmt.Sprintf(format+"\r\n", args...))
	conn.writer.Flush()
}

// initialize initializes the connection context
func (conn *FtpConnection) getDataConnection() (net.Conn, error) {
	if conn.passive {
		conn.passive = false
		defer conn.listener.Close()

		// Manage incoming connection
		cnx, err := conn.listener.Accept()
		if err != nil {
			log.Printf("commandListPassive: Unable to accept incoming connection: %v", err)
			conn.writeString("425 Can't open data connection.")
			return nil, fmt.Errorf("unable to accept incoming connection: %v", err)
		}
		return cnx, nil
	} else {
		remoteAddr := conn.remoteDataEndPoint.String()
		if remoteAddr == "" {
			log.Printf("commandListActive: no remote data endpoint")
			conn.writeString("451 Requested action aborted: local error in processing.")
			return nil, fmt.Errorf("no remote data endpoint")
		}

		// Establish the data connection
		cnx, err := net.Dial("tcp", remoteAddr)
		if err != nil {
			log.Printf("commandListActive: unable to establish data connection: %s (%v)", remoteAddr, err)
			conn.writeString("425 Can't open data connection.")
			return nil, fmt.Errorf("unable to establish connection: %v", err)
		}
		return cnx, nil
	}
}

// commandUser manages the USER FTP command
func (conn *FtpConnection) commandUser(args []string) {
	// Check arguments count
	if len(args) != 1 {
		log.Printf("commandUser: bad arguments count: %v", args)
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	// Store the user name
	conn.user = args[0]

	// Clear the authentication
	conn.authenticated = false

	// Send the reply
	conn.writeString("331 User name ok, need password.")
}

// commandPassword manages the PASS FTP command
func (conn *FtpConnection) commandPassword(args []string) {
	// Check arguments count
	if len(args) != 1 {
		log.Printf("commandPassword: bad arguments count: %v", args)
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	// Check a user name is defined
	if conn.user == "" {
		conn.writeString("332 Need account for login.")
		return
	}

	//if conn.user != "anonymous" {
	// TODO : Check password
	//}
	conn.authenticated = true

	// Send the reply
	conn.writeString("230 User logged in, proceed.")
}

// commandSystem manages the SYST FTP command
func (conn *FtpConnection) commandSystem(args []string) {
	// Check arguments count
	if len(args) != 0 {
		log.Printf("commandSystem: bad arguments count: %v", args)
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	// Send the reply
	conn.writeString("215 UNIX")
}

// commandPassive manages the PASV FTP command
func (conn *FtpConnection) commandPassive(args []string) {
	// Check arguments count
	if len(args) != 0 {
		log.Printf("commandPassive: bad arguments count: %v", args)
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	// Get the local address of the connection with remote
	localAddress := conn.tcpConnection.LocalAddr().String()

	// Get the host from the local address
	localHost, _, err := net.SplitHostPort(localAddress)
	if err != nil {
		log.Printf("commandPassive: unable to parse local address: %s (%v)", localAddress, err)
		conn.writeString("451 Requested action aborted: local error in processing.")
		return
	}

	// Create the listener (the port is dynamically allocated)
	listener, err := net.Listen("tcp4", net.JoinHostPort(localHost, ""))
	if err != nil {
		log.Printf("commandPassive: unable to create listener: %v", err)
		conn.writeString("451 Requested action aborted: local error in processing.")
		return
	}

	// Get the listener address
	listenerAddr := listener.Addr().String()

	// Parse the listener address
	listenerTCPAddr, err := net.ResolveTCPAddr("tcp", listenerAddr)
	if err != nil {
		log.Printf("commandPassive: unable to parse endpoint: %s (%v)", listenerAddr, err)
		conn.writeString("451 Requested action aborted: local error in processing.")
		return
	}

	// Get the IP V4 address of the listener
	ipv4 := listenerTCPAddr.IP.To4()
	if ipv4 == nil {
		log.Printf("commandPassive: listener is not an IP V4 endpoint: %s", listenerAddr)
		conn.writeString("451 Requested action aborted: local error in processing.")
		return
	}

	// Send the reply
	conn.writeString("227 Entering passive mode (%d,%d,%d,%d,%d,%d).",
		ipv4[0], ipv4[1], ipv4[2], ipv4[3],
		listenerTCPAddr.Port/256, listenerTCPAddr.Port%256)

	conn.listener = listener
	conn.passive = true
}

// commandExtendedPassive manages the EPSV FTP command
func (conn *FtpConnection) commandExtendedPassive(args []string) {
	// Check arguments count
	if len(args) != 0 {
		log.Printf("commandExtendedPassive: bad arguments count: %v", args)
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	// Get the local address of the connection with remote
	localAddress := conn.tcpConnection.LocalAddr().String()

	// Get the host from the local address
	localHost, _, err := net.SplitHostPort(localAddress)
	if err != nil {
		log.Printf("commandPassive: unable to parse local address: %s (%v)", localAddress, err)
		conn.writeString("451 Requested action aborted: local error in processing.")
		return
	}

	// Create the listener (the port is dynamically allocated)
	listener, err := net.Listen("tcp", net.JoinHostPort(localHost, ""))
	if err != nil {
		log.Printf("commandExtendedPassive: unable to create listener: %v", err)
		conn.writeString("451 Requested action aborted: local error in processing.")
		return
	}

	// Get listener address
	listenerAddr := listener.Addr().String()

	// Get the port from the listener address
	_, port, err := net.SplitHostPort(listenerAddr)
	if err != nil {
		log.Printf("commandExtendedPassive: unable to parse listener address: %s (%v)", listenerAddr, err)
		conn.writeString("451 Requested action aborted: local error in processing.")
		return
	}

	// Send the reply
	conn.writeString("229 Entering extended passive mode (|||%s|).", port)

	conn.listener = listener
	conn.passive = true
}

// commandChangeWorkingDirectory manages the CWD FTP command
func (conn *FtpConnection) commandChangeWorkingDirectory(args []string) {
	// Check arguments count
	if len(args) != 1 {
		log.Printf("commandChangeWorkingDirectory: bad arguments count: %v", args)
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	// Check the new value
	if args[0] == "" {
		log.Printf("commandChangeWorkingDirectory: no new working directory")
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	// Compute the new working directory
	var tmpPath string
	if args[0][0] == '/' {
		tmpPath = args[0]
	} else {
		tmpPath = filepath.Join(conn.workingDirectory, args[0])
	}

	// Check the resulting value
	if tmpPath == "" {
		log.Printf("commandChangeWorkingDirectory: resulting new working directory is empty")
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	stat, err := os.Stat(tmpPath)
	if err != nil {
		log.Printf("commandChangeWorkingDirectory: new working directory does not exist: %s", tmpPath)
		conn.writeString("431 No such directory")
		return
	}

	if !stat.IsDir() {
		log.Printf("commandChangeWorkingDirectory: new working directory is not a directory: %s", tmpPath)
		conn.writeString("431 No such directory")
		return
	}

	// Store the working directory
	conn.workingDirectory = tmpPath

	// Send the reply
	conn.writeString("200 directory changed to %s.", conn.workingDirectory)
}

// commandPrintWorkingDirectory manages the PWD FTP command
func (conn *FtpConnection) commandPrintWorkingDirectory(args []string) {
	// Check arguments count
	if len(args) != 0 {
		log.Printf("commandPrintWorkingDirectory: bad arguments count: %v", args)
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	// Send the reply
	conn.writeString("200 working directory is %s.", conn.workingDirectory)
}

// commandList manages the LIST FTP command
func (conn *FtpConnection) commandList(args []string) {
	// Check arguments count
	if len(args) != 0 {
		log.Printf("commandList: bad arguments count: %v", args)
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	// Get the directory list
	dirs, err := ioutil.ReadDir(conn.workingDirectory)
	if err != nil {
		log.Printf("commandList: unable to read directory: %v", err)
		conn.writeString("553 Requested action not taken.")
		return
	}

	// Send the preliminary reply
	conn.writeString("150 Here comes the directory listing.")

	// Get the data connection
	dataConnection, err := conn.getDataConnection()
	if err != nil {
		log.Printf("commandList: Unable to get data connection: %v", err)
		return
	}
	defer dataConnection.Close()

	// Send the directory list
	for _, dir := range dirs {
		line := fmt.Sprintf("%s %10d %15s %s\r\n", dir.Mode().String(), dir.Size(), dir.ModTime().Format("2 Jan 2006"), dir.Name())
		written, err := dataConnection.Write([]byte(line))
		if err != nil {
			log.Printf("commandList: Unable to send data: %v", err)
			conn.writeString("451 Requested action aborted: local error in processing.")
			return
		}
		if written != len(line) {
			log.Printf("commandList: Unable to send all data: %d/%d", written, len(line))
			conn.writeString("451 Requested action aborted: local error in processing.")
			return
		}
	}

	// Send the completion reply
	conn.writeString("226 Directory send OK.")
}

// commandStorage manages the STOR FTP command
func (conn *FtpConnection) commandStorage(args []string) {
	// Check arguments count
	if len(args) != 1 {
		log.Printf("commandStorage: bad arguments count: %v", args)
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	// Check a user name is defined
	if !conn.authenticated {
		log.Printf("commandStorage: not logged in")
		conn.writeString("530 Not logged in.")
		return
	}

	// Send the preliminary reply
	conn.writeString("150 File status okay; about to open data connection.")

	// Establish the data connection
	dataConnection, err := conn.getDataConnection()
	if err != nil {
		log.Printf("commandStorage: unable to get data connection: %v", err)
		return
	}
	defer dataConnection.Close()

	// Receive the data
	var data []byte
	for {
		tmp := make([]byte, 4096)
		read, err := dataConnection.Read(tmp)
		if err != nil {
			break
		}
		if read == 0 {
			break
		}
		data = append(data, tmp...)
	}

	// Convert data into local format
	if !conn.binary {
		var tmp []byte
		for _, b := range data {
			switch b {
			case '\r':
			default:
				tmp = append(tmp, b)
			}
		}
		data = tmp
	}

	// Store the data into the file
	fullName := filepath.Join(conn.workingDirectory, args[0])
	err = ioutil.WriteFile(fullName, data, 0777)
	if err != nil {
		log.Printf("commandStorage: unable to store data: %s (%v)", fullName, err)
		conn.writeString("552 Requested file action aborted.")
		return
	}

	// Send the completion reply
	conn.writeString("226 Closing data connection, file transfer successful.")
}

// commandRetreive manages the RETR FTP command
func (conn *FtpConnection) commandRetrieve(args []string) {
	// Check arguments count
	if len(args) != 1 {
		log.Printf("commandRetrieve: bad arguments count: %v", args)
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	// Check a user name is defined
	if !conn.authenticated {
		log.Printf("commandRetrieve: not logged in")
		conn.writeString("530 Not logged in.")
		return
	}

	// Read the file
	fullName := filepath.Join(conn.workingDirectory, args[0])
	data, err := ioutil.ReadFile(fullName)
	if err != nil {
		log.Printf("commandRetrieve: unable to read file: %s (%v)", fullName, err)
		conn.writeString("550 Requested action not taken. File not found.")
		return
	}

	// Convert data into network format
	if !conn.binary {
		var tmp []byte
		var prevCR bool
		var prevCRLF bool
		for _, b := range data {
			switch b {
			case '\n':
				if !prevCR {
					tmp = append(tmp, '\r')
				}
				tmp = append(tmp, b)
				prevCR = false
				prevCRLF = true
			case '\r':
				prevCR = true
				prevCRLF = false
				tmp = append(tmp, b)
			default:
				if prevCR && !prevCRLF {
					tmp = append(tmp, '\n')
				}
				prevCR = false
				prevCRLF = false
				tmp = append(tmp, b)
			}
		}
		data = tmp
	}

	// Send the preliminary reply
	conn.writeString("150 File status okay; about to open data connection.")

	// Establish the data connection
	dataConnection, err := conn.getDataConnection()
	if err != nil {
		log.Printf("commandRetreive: unable to get data connection: %v", err)
		return
	}
	defer dataConnection.Close()

	// Send data
	written, err := dataConnection.Write(data)
	if err != nil {
		log.Printf("commandRetrieve: unable to send data: %v", err)
		conn.writeString("426 Connection closed; transfer aborted.")
		return
	}

	if written != len(data) {
		log.Printf("commandRetrieve: Unable to send all data: %d/%d", written, len(data))
		conn.writeString("426 Connection closed; transfer aborted.")
		return
	}

	// Send the completion reply
	conn.writeString("226 Closing data connection, file transfer successful.")
}

// commandPort manages the PORT FTP command
func (conn *FtpConnection) commandPort(args []string) {
	// Check arguments count
	if len(args) != 1 {
		log.Printf("commandPort: bad arguments count: %v", args)
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	// Split parameters
	params := strings.Split(args[0], ",")
	if len(params) != 6 {
		log.Printf("commandPort: bad bytes count: %v", params)
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	// Parse each argument
	var a []byte = make([]byte, 6)
	for i, param := range params {
		// Convert the argument as an integer
		v, err := strconv.ParseInt(param, 10, 32)
		if err != nil {
			log.Printf("commandPort: bad integer:  %s", param)
			conn.writeString("501 Syntax error in parameters or arguments.")
			return
		}
		// Check the value range
		if v < 0 || v > 255 {
			log.Printf("commandPort: invalid integer: %d", v)
			conn.writeString("501 Syntax error in parameters or arguments.")
			return
		}

		// Assign the value to the resulting array
		a[i] = byte(v)
	}

	// Compute the port number
	port := int(a[4])*256 + int(a[5])

	// Check the value range
	if port < 1 || port > 65535 {
		log.Printf("commandPort: invalid port:  %d", port)
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	// Store the new values
	conn.remoteDataEndPoint.IP = conn.remoteDataEndPoint.IP[:4]
	for i := 0; i < 4; i++ {
		conn.remoteDataEndPoint.IP[i] = a[i]
	}
	conn.remoteDataEndPoint.Port = port

	// Send the reply
	conn.writeString("200 Command okay.")
}

// commandExtendedPort manages the EPRT FTP command
func (conn *FtpConnection) commandExtendedPort(args []string) {
	// Check arguments count
	if len(args) != 1 {
		log.Printf("commandExtendedPort: bad arguments count: %v", args)
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	// Parse the argument
	delim := args[0][0]
	params := strings.Split(args[0], string(delim))
	if len(params) != 5 {
		log.Printf("commandExtendedPort: bad parameters count: %v", params)
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	// Build the remote address in standard format
	remoteAddr := net.JoinHostPort(params[2], params[3])

	// Parse the remote address
	tcpAddr, err := net.ResolveTCPAddr("tcp", remoteAddr)
	if err != nil {
		log.Printf("commandExtendedPort: invalid address: %s (%v)", remoteAddr, err)
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	// Store the remote address
	conn.remoteDataEndPoint = *tcpAddr

	// Send the reply
	conn.writeString("200 Command okay.")
}

// commandType manages the TYPE FTP command
func (conn *FtpConnection) commandType(args []string) {
	// Check arguments count
	if len(args) < 1 {
		log.Printf("commandType: bad arguments count: %v", args)
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	switch args[0] {
	case "A": // ASCII
		switch len(args) {
		case 1:
			conn.writeString("200 Command okay.")
			conn.binary = false
		case 2:
			switch args[1] {
			case "N":
				conn.writeString("200 Command okay.")
			case "T":
				conn.writeString("504 Command not implemented for that parameter.")
			case "C":
				conn.writeString("504 Command not implemented for that parameter.")
			default:
				log.Printf("commandType: unknown sub-type: %s", args[1])
				conn.writeString("501 Syntax error in parameters or arguments.")
			}
		default:
			log.Printf("commandType: invalid arguments count: %v", args)
			conn.writeString("501 Syntax error in parameters or arguments.")
		}
	case "E": // EBCDIC
		conn.writeString("504 Command not implemented for that parameter.")
	case "I": // IMAGE
		conn.writeString("200 Command okay.")
		conn.binary = true
	case "L": // LOCAL
		conn.writeString("504 Command not implemented for that parameter.")
	default:
		log.Printf("commandType: unknown type: %s", args[0])
		conn.writeString("501 Syntax error in parameters or arguments.")
	}
}

// commandStructure manages the STRU FTP command
func (conn *FtpConnection) commandStructure(args []string) {
	// Check arguments count
	if len(args) != 1 {
		log.Printf("commandStructure: bad arguments count: %v", args)
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	switch args[0] {
	case "F": // File (no record structure)
		conn.writeString("200 Command okay.")
	case "R": // Record structure
		conn.writeString("504 Command not implemented for that parameter.")
	case "P": // Page structure
		conn.writeString("504 Command not implemented for that parameter.")
	default:
		log.Printf("commandStructure: unknown structure: %s", args[0])
		conn.writeString("501 Syntax error in parameters or arguments.")
	}
}

// commandMode manages the MODE FTP command
func (conn *FtpConnection) commandMode(args []string) {
	// Check arguments count
	if len(args) != 1 {
		log.Printf("commandMode: bad arguments count: %v", args)
		conn.writeString("501 Syntax error in parameters or arguments.")
		return
	}

	switch args[0] {
	case "S": // Stream
		conn.writeString("200 Command okay.")
	case "B": // Block
		conn.writeString("504 Command not implemented for that parameter.")
	case "C": // Compressed
		conn.writeString("504 Command not implemented for that parameter.")
	default:
		log.Printf("commandMode: unknown mode: %s", args[0])
		conn.writeString("501 Syntax error in parameters or arguments.")
	}
}

// handle manages a ftp connection
func (conn *FtpConnection) handle() {
	defer conn.tcpConnection.Close()

	// Initialize the connection context
	err := conn.initialize()
	if err != nil {
		log.Printf("Unable to initialize context: %v", err)
		return
	}

	// Create the input scanner
	scanner := bufio.NewScanner(conn.tcpConnection)

	conn.writeString("220 Service ready.")

	for scanner.Scan() {

		// Get the command line
		line := scanner.Text()

		log.Printf("CMD: %s\n", line)

		// Split the command line
		tokens := strings.Split(line, " ")
		if len(tokens) == 0 {
			continue
		}

		// Dispatch according to the command
		switch strings.ToUpper(tokens[0]) {
		case "USER":
			conn.commandUser(tokens[1:])
		case "PASS":
			conn.commandPassword(tokens[1:])
		case "SYST":
			conn.commandSystem(tokens[1:])
		case "PASV":
			conn.commandPassive(tokens[1:])
		case "EPSV":
			conn.commandExtendedPassive(tokens[1:])
		case "CWD":
			conn.commandChangeWorkingDirectory(tokens[1:])
		case "PWD":
			conn.commandPrintWorkingDirectory(tokens[1:])
		case "LIST":
			conn.commandList(tokens[1:])
		case "STOR":
			conn.commandStorage(tokens[1:])
		case "RETR":
			conn.commandRetrieve(tokens[1:])
		case "PORT":
			conn.commandPort(tokens[1:])
		case "EPRT":
			conn.commandExtendedPort(tokens[1:])
		case "TYPE":
			conn.commandType(tokens[1:])
		case "STRU":
			conn.commandStructure(tokens[1:])
		case "MODE":
			conn.commandMode(tokens[1:])
		case "QUIT":
			conn.writeString("221 Service closing control connection. Logged out if appropriate.")
			return
		case "NOOP":
			conn.writeString("200 Command okay.")
		default:
			log.Printf("Unknown command: %s", line)
			conn.writeString("502 Command not implemented.")
		}
	}
}

// main is the entry point of the program
func main() {
	// Get parameters
	port := flag.Int("port", 21, "Listen port")
	flag.Parse()

	if *port < 1 || *port > 65535 {
		log.Fatalf("Invalid port: %d", *port)
	}

	// Listen incoming connections
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Unable to listen on port %d: %v", port, err)
	}
	defer listener.Close()

	// Manage incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Unable to accept incoming connection: %v", err)
			continue
		}
		ftpConnection := &FtpConnection{tcpConnection: conn}
		go ftpConnection.handle()
	}
}
