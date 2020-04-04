package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("No port number specified... Exiting....")
		return
	}
	addr := arguments[1]
	// Connects to server
	connection, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Please use format 'Address:Port' or ':Port' for localhost")
		return

	}
	for {
		// Client -> Server
		fmt.Print("Text to send: ")
		inputscanner := bufio.NewScanner(os.Stdin)
		inputscanner.Scan()
		text := inputscanner.Text()
		fmt.Fprintf(connection, text+"\n")

		// Server -> Client
		replyscanner, error := bufio.NewReader(connection).ReadString('\n')
		if error != nil {
			fmt.Println("ERROR : Cannot reach server")
			return
		}
		fmt.Println(replyscanner)
	}
}
