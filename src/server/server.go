package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	// Listen for new connections;
	listener, err := net.Listen("tcp", HOST + ":" + PORT)
	if err != nil {
		log.Fatal("ERROR: Could not listen on PORT %s", PORT)
	}
	defer listener.Close()

	fmt.Println("Listening on ", HOST + ":" + PORT)
	
	// Listen and accept new connections;
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("ERROR: Could not accept connection %d", err.Error())
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
	fmt.Println("Client connected on %s", conn.RemoteAddr().String())

	// Handle incoming commands from the clients;
	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected or errored out %s", err.Error())
			return
		}
		// Print the message from the client;
		fmt.Println("Client: %s", msg)

		// Echo the msg back to the client;
		resp := []byte("MSG recieved: " + msg)
		conn.Write(resp)
	}
}
