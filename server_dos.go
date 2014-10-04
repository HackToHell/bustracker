// server_dos
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	for {
		go Spamreqs()
	}
}

func Spamreqs() {
	l, err := net.Dial("tcp", "localhost:3334")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Hello World!")
	// Close the listener when the application closes.
	l.Close()
}
