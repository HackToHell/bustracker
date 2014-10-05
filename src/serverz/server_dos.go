// server_dos
package serverz

import (
	"fmt"
	"net"
	"os"
)

func main() {

	go Spamreqs()

}

func Spamreqs() {
	l, err := net.Dial("tcp", "localhost:8090")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Hello World!")
	// Close the listener when the application closes.
	l.Close()
}
