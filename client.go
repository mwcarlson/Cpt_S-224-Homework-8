package main

import 
(
	"fmt"
	"net"
	"bufio"
)

func main() {
	// connect to 'localhost' on port '7778'
	con, err := net.Dial("tcp", "localhost:7778")
	if err != nil {
		panic(err)
	}

	// write a message
	fmt.Fprintf(con, "Hello Internet!")
	
	// read out the information sent to the screen 
	scan := bufio.NewScanner(con)
	for scan.Scan() {
		fmt.Println(scan.Text())	
	}

	// close the connection
	con.Close()
}

