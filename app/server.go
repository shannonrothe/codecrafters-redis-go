package main

import (
	"fmt"
	"io"
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
		b := make([]byte, 1024)
		if _, err := c.Read(b); err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("error reading from client: ", err.Error())
				os.Exit(1)
			}
		}
		c.Write([]byte("+PONG\r\n"))
	}
}
