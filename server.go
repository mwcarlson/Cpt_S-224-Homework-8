package main

import 
(
	"io"
	"net"
	"os"
)

func handleConnection(c net.Conn) 
{
	// Copy up to 128 bytes from the connection to the screen.
	io.CopyN(os.Stdout, c, 128)
	for range time.Tick(time.Second)
	{
		fmt.Fprintf(con, "%s\n", time.Now().String())
	}
}

func main() 
{
	// listen for connections on port 5555
	list, err := net.Listen("tcp", ":7778")
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
