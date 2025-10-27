package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"bufio"
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
	fmt.Println("Please enter a message to send")

	// Init a new reader to scan from stdin;
	reader := bufio.NewReader(os.Stdin)

	// Until connection is closed keep reading the users messages;
	for {
		// Collect string from the reader;
		fmt.Printf(">: ")
		text, _ := reader.ReadString('\n')

		// Handle graceful exit if text = "EXIT";
		if (strings.TrimSpace(text) == "EXIT") {
			conn.Close()
			return
		}
		
		// Send the string to the server as bytes;
		_, err := conn.Write([]byte(text))
		if (err != nil) {
			fmt.Printf("ERROR: could not send message %s", err.Error())
			return
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
