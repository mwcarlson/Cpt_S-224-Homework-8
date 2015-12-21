package main

import (
	"io"
	"net"
	"os"
)

func handleConnection(c net.Conn) 
{
	// Copy up to 128 bytes from the connection to the screen.
	io.CopyN(os.Stdout, c, 128)
}

func main() 
{
	// listen for connections on port 5555
	list, err := net.Listen("tcp", "localhost:7778")
	if err != nil 
	{
		// If others are using the same computer
		// you might want to change the port
		panic(err)
	}

	// loop forever!
	for 
	{
		con, err := list.Accept()
		if err != nil 
		{
			panic(err)
		}

		go handleConnection(con)
	}
}
