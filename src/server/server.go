package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"log"
	"net"
)

const (
	ADDR = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	// Listen for new connections;
	listener, err := net.Listen("tcp", ADDR + ":" + PORT)
	if (err != nil) {
		log.Fatalf("ERROR: Could not listen on PORT %s", PORT)
	}
	defer listener.Close()

	fmt.Printf("Listening on %s:%s\n", ADDR, PORT)
	
	// Listen and accept new connections;
	for {
		conn, err := listener.Accept()
		if (err != nil) {
			log.Printf("ERROR: Could not accept connection %s", err.Error())
			continue
		}
		// Handle the new conneciton;
		go HandleConnection(conn)
	}
}

// Handles the logic for a new connection;
func HandleConnection(conn net.Conn) {
	// defer connection closing;

	// Prints the IP of the connection;
	fmt.Printf("Client connected on %s\n", conn.RemoteAddr().String())

	// Handle incoming commands from the clients;
	for {
		// Create the file;
		file, err := os.Create("./cat")
		if (err != nil) {
			fmt.Printf("ERROR: Could not create file: %s\n", err.Error())
		}

		// Create buffer for reading file from conn;
		buf := make([]byte, 8)
		// Create reader to read bytes into buf;
		reader := bufio.NewReader(conn)
		for {
			// Read data from conn and write it to the file;
			n, err := io.ReadFull(reader, buf)
			if n > 0 {
				// Write buffer to file;
				_, err = file.Write(buf)
				if (err != nil) {
					fmt.Printf("ERROR: Problem writing buf to file %s", err.Error())
					return
				}
			}

			if (err == io.EOF) {
				fmt.Printf("Reached EOF\n")
				return
			} 

			if (err == io.ErrUnexpectedEOF) {
				fmt.Printf("ERROR: Problem reading file: %s\n", err.Error())
				break
			}
		}

		// Echo the msg back to the client;
		resp := []byte("File recieved: ")
		conn.Write(resp)
	}
}
