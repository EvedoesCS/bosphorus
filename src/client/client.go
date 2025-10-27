package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

const (
	// server addr;
	ADDR = "127.0.0.1:8080"
)

func main() {
	// Attempt to dial to the server; 
	conn, err := net.Dial("tcp", ADDR)
	if (err != nil) {
		log.Fatalf("ERROR: Could not dial server %s", err.Error())
	}
	defer conn.Close()

	fmt.Printf("SUCCESS: Connected to the server %s", ADDR)
	fmt.Println("Please enter the path of the file to send")

	// Init a new reader to scan from stdin;
	reader := bufio.NewReader(os.Stdin)

	// Until connection is closed keep reading the users messages;
	for {
		// Collect string from the reader;
		fmt.Printf(">: ")
		filepath, _ := reader.ReadString('\n')
		filepath = strings.TrimSpace(filepath)
		filepath = "./" + filepath

		// Handle graceful exit if text = "EXIT";
		if (strings.TrimSpace(filepath) == "EXIT") {
			conn.Close()
			return
		}

		// Read the file in buffered chunks
		buf := make([]byte, 8)

		file, err := os.Open(filepath)
		if (err != nil) {
			fmt.Printf("ERROR: Could not open file %s for reading: %s", filepath, err.Error())
			return
		}
		defer file.Close()

		// Read the file in chunks of sizeof(buf);
		reader := bufio.NewReader(file)
		for {
			n, err := io.ReadFull(reader, buf)

			
			// Send the buf through conn;
			if n > 0 {
				_, err = conn.Write([]byte(buf))
				if (err != nil) {
					fmt.Printf("ERROR: could not send message %s", err.Error())
					return
				}
			}
			// Break when EOF reached;
			if (err == io.EOF) {
				return	
			}

			// Break if a non-EOF error occurs;
			if (err == io.ErrUnexpectedEOF) {
				fmt.Printf("ERROR: Problem reading file: %s\n", err.Error())
				break
			}

		}

		// Wait for a response;
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if (err != nil) {
			fmt.Printf("ERROR: Could not read response from the server %s", err.Error())
			return 
		}

		// Print the msg from the server;
		fmt.Printf("SERVER: %s", msg)
	}

}
