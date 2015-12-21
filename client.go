package main

import 
(
	"fmt"
	"net"
)

func main() 
{
	// connect to 'localhost' on port '7778'
	con, err := net.Dial("tcp", "localhost:7778")
	if err != nil 
	{
		panic(err)
	}

	// Write a message
	fmt.Fprintf(con, "Hello Internet!")

	// Close the connection
	con.Close()
}

