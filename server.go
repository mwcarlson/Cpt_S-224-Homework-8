package main

import 
(
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func handleConnection(c net.Conn) 
{
	// Copy up to 128 bytes from the connection to the screen.
	io.CopyN(os.Stdout, c, 128)
	fmt.Fprintf(c, "The time is %s\n", time.Now().String())
}

func main() 
{
	
}
