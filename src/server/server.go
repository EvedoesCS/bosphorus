package main 

import (
	"bufio"
	"fmt"
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
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if (err != nil) {
			fmt.Printf("Client disconnected or errored out %s\n", err.Error())
			return
		}
		// Print the message from the client;
		fmt.Printf("Client: %s\n", msg)

		// Echo the msg back to the client;
		resp := []byte("MSG recieved: " + msg)
		conn.Write(resp)
	}
}
