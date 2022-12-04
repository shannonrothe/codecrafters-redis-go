package main

import (
	"fmt"
	"net"
	"os"
	// Uncomment this block to pass the first stage
	// "net"
	// "os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	c, err := l.Accept()
	defer c.Close()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	for {
		b := new([]byte)
		n, err := c.Read(*b)
		fmt.Printf("Reading from connection... %d bytes\n", n)
		if err != nil {
			fmt.Println("Failed to read bytes", err.Error())
			os.Exit(1)
		}
		c.Write([]byte("+PONG\r\n"))
	}
}
